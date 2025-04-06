import { AuthContext } from '@/app/(auth)/AuthContext'
import EventCard from '@/components/EventCard'
import { colors } from '@/constants/colors'
import { fonts } from '@/constants/fonts'
import AsyncStorage from '@react-native-async-storage/async-storage'
import { Link, router } from 'expo-router'
import { useCallback, useContext, useEffect, useState } from 'react'
import {
  Button,
  Image,
  RefreshControl,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native'
import { SafeAreaView } from 'react-native-safe-area-context'
import { format, parseISO } from 'date-fns'
import { ru, tr } from 'date-fns/locale'
import { IP_ADDRESS, PORT } from '@/constants/address'

type Event = {
  ID: number
  title: string
  description: string
  category: string
  is_public: boolean
  status: string
  date: string
  start_time: string
  end_time: string
  location: string
  creator_id: number
  // organization_id: number;
}

export default function Home() {
  const auth = useContext(AuthContext)

  const [joinedEvents, setJoinedEvents] = useState<Event[]>([])
  const [openEvents, setOpenEvents] = useState<Event[]>([])

  const [loading, setLoading] = useState(false)

  const [refreshing, setRefreshing] = useState(false)

  const formatDate = (date: string) => {
    const parsedDate = parseISO(date)
    const formattedDate = format(parsedDate, 'd MMMM', {
      locale: ru,
    }).toUpperCase()

    return formattedDate
  }

  const formatTime = (time: string) => {
    if (!time) return '00:00'
    return time.slice(0, 5)
  }

  const LoadEvents = async () => {
    setRefreshing(true)
    const timeout = setTimeout(() => setLoading(true), 300)
    try {
      const token = await AsyncStorage.getItem('token')
      if (!token) throw new Error('Отсутствует токен')

      const response = await fetch(IP_ADDRESS + PORT + '/getevents', {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      })

      if (!response.ok) {
        if (response.status === 401) {
          router.replace('/(auth)/login')
        }
        if (response.status === 404) {
          setJoinedEvents([])
          setOpenEvents([])
          return
        }
        throw new Error(`Ошибка запроса: ${response.status}`)
      }

      const data = await response.json()

      const joined = Array.isArray(data.joined_events) ? data.joined_events : []

      const open = Array.isArray(data.open_events) ? data.open_events : []

      setJoinedEvents(joined)

      const joinedEventIds = joined.map((event: Event) => event.ID)

      const filteredOpenEvents = open.filter(
        (event: Event) => !joinedEventIds.includes(event.ID)
      )

      setOpenEvents(filteredOpenEvents)
    } catch (error) {
      console.error('Ошибка запроса:', error)
    } finally {
      clearTimeout(timeout)
      setLoading(false)
      setRefreshing(false)
    }
  }

  useEffect(() => {
    LoadEvents()
    const interval = setInterval(() => {
      LoadEvents()
    }, 30000)

    return () => clearInterval(interval)
  }, [])

  const renderJoinedEvents = () => {
    return joinedEvents.map((event) => {
      return (
        <View style={styles.category__item} key={'activeJoined_' + event.ID}>
          <EventCard
            id={event.ID}
            title={event.title}
            description={event.description}
            category={event.category}
            is_public={event.is_public}
            status={event.status}
            date={formatDate(event.date)}
            start_time={formatTime(event.start_time)}
            end_time={formatTime(event.end_time)}
            location={event.location}
            creator_id={event.creator_id}
            joined={true}
          />
        </View>
      )
    })
  }

  const renderOpenEvents = () => {
    return openEvents.map((event) => {
      return (
        <View style={styles.category__item} key={'active_' + event.ID}>
          <EventCard
            id={event.ID}
            title={event.title}
            description={event.description}
            category={event.category}
            is_public={event.is_public}
            status={event.status}
            date={formatDate(event.date)}
            start_time={formatTime(event.start_time)}
            end_time={formatTime(event.end_time)}
            location={event.location}
            creator_id={event.creator_id}
            joined={false}
          />
        </View>
      )
    })
  }

  return (
    <SafeAreaView style={{ backgroundColor: '#ffffff' }}>
      <View style={styles.header}>
        <View style={styles.header__content}>
          <Image
            source={require('../../../assets/icons/logo.png')}
            style={{ width: 20, height: 20 }}
            resizeMode="contain"
          />
          <Text style={styles.header__title}>EVENTHUB</Text>
        </View>
      </View>
      <ScrollView
        style={styles.categories}
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={LoadEvents} />
        }
      >
        {joinedEvents.length > 0 ? (
          <View style={styles.category}>
            <Text style={styles.category__title}>Вы участвуете</Text>
            <ScrollView
              horizontal={true}
              showsHorizontalScrollIndicator={false}
              style={styles.category__content}
            >
              {renderJoinedEvents()}
            </ScrollView>
          </View>
        ) : null}

        {openEvents.length > 0 ? (
          <View style={styles.category}>
            <Text style={styles.category__title}>Открытые мероприятия</Text>
            <ScrollView
              horizontal={true}
              showsHorizontalScrollIndicator={false}
              style={styles.category__content}
            >
              {renderOpenEvents()}
            </ScrollView>
          </View>
        ) : null}
      </ScrollView>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  header: {
    position: 'fixed',
    height: 46,
    display: 'flex',
    alignItems: 'center',
    paddingTop: 4,
    boxSizing: 'border-box',
  },
  header__content: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignContent: 'center',
  },
  header__title: {
    fontFamily: fonts.Unbounded,
    fontWeight: 800,
    color: colors.primary,
    marginLeft: 8,
    fontSize: 16,
  },
  categories: {
    width: '100%',
    height: '100%',
  },
  category: {
    marginBottom: 32,
  },
  category__title: {
    color: colors.grey_text,
    fontFamily: fonts.Unbounded,
    fontSize: 16,
    fontWeight: 700,
    marginLeft: 16,
    marginBottom: 16,
  },
  category__content: {
    paddingLeft: 16,
  },
  category__item: {
    marginRight: 12,
  },
})
