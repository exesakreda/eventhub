import {
  View,
  Text,
  Button,
  Platform,
  Image,
  StyleSheet,
  TouchableOpacity,
  ScrollView,
  Alert,
} from 'react-native'
import { StatusBar } from 'expo-status-bar'
import { useLocalSearchParams, useRouter } from 'expo-router'
import { fonts } from '@/constants/fonts'
import { colors } from '@/constants/colors'
import AsyncStorage from '@react-native-async-storage/async-storage'
import * as Haptics from 'expo-haptics'
import CustomButton from '@/components/CustomButton'
import { IP_ADDRESS, PORT } from '@/constants/address'

const ModalScreen = () => {
  const router = useRouter()

  const {
    id,
    is_public,
    title,
    category,
    description,
    date,
    start_time,
    end_time,
    location,
    creator_id,
    joined,
  } = useLocalSearchParams()

  const JoinEvent = async (event_id: number) => {
    try {
      const token = await AsyncStorage.getItem('token')
      if (!token) throw new Error('Отсутствует токен')

      const response = await fetch(
        'http://' + IP_ADDRESS + ':' + PORT + '/join_event',
        {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ event_id }),
        }
      )

      if (!response.ok) {
        if (response.status === 401) {
          router.replace('/(auth)/login')
        }
        throw new Error(`Ошибка запроса: ${response.status}`)
      }
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Success)
    } catch (error) {
      Alert.alert('Ошибка сервера', 'Проверьте подключение.')
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error)
    }
  }

  const QuitEvent = async (event_id: number) => {
    try {
      const token = await AsyncStorage.getItem('token')
      if (!token) throw new Error('Отсутствует токен')

      const response = await fetch(
        'http://' + IP_ADDRESS + ':' + PORT + '/quit_event',
        {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ event_id }),
        }
      )

      if (!response.ok) {
        const errorData = await response.json()
        if (response.status === 401) {
          router.replace('/(auth)/login')
        }
        throw new Error(`Ошибка запроса: ${response.status}`)
      }
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Success)
    } catch (error) {
      Alert.alert('Ошибка сервера', 'Проверьте подключение.')
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error)
    }
  }

  return (
    <View style={{ flex: 1, backgroundColor: '#ffffff' }}>
      <TouchableOpacity onPress={() => router.back()} style={styles.close}>
        <Image
          source={require('../assets/icons/close.png')}
          style={{ height: 16, width: 16 }}
        />
      </TouchableOpacity>
      <ScrollView style={styles.container}>
        <Image
          source={require('../assets/images/blank_image_modal.png')}
          style={{ height: '50%', width: '100%' }}
        />

        <View style={styles.content}>
          <View style={styles.tags}>
            {joined == 'true' ? (
              <View style={styles.join}>
                <Text style={styles.join__text}>ВЫ УЧАСТВУЕТЕ</Text>
              </View>
            ) : null}
            <View style={styles.public}>
              <Text style={styles.public__text}>
                {is_public ? 'ОТКРЫТОЕ' : 'ПО ПРИГЛАШЕНИЯМ'}
              </Text>
            </View>
          </View>
          <View style={styles.mainText}>
            <Text style={styles.title}>{title}</Text>
            <Text style={styles.category}>{category}</Text>
            <Text style={styles.description}>{description}</Text>
          </View>

          <View style={styles.date}>
            <View style={styles.date__item}>
              <Text style={styles.item__text}>{date}</Text>
            </View>
            <View style={styles.date__item}>
              <Text style={styles.item__text}>
                {start_time} - {end_time}
              </Text>
            </View>
          </View>
          <View style={styles.location}>
            <Image
              source={require('../assets/icons/location.png')}
              style={{ height: 18, width: 14 }}
            />
            <Text style={styles.location__text}>{location}</Text>
          </View>
        </View>

        <Text>{creator_id}</Text>
      </ScrollView>

      <View style={styles.button}>
        {joined == 'true' ? (
          <CustomButton
            onPress={() => QuitEvent(Number(id))}
            title="Отменить участие"
            type={'cancel'}
            fill={'bordered'}
          />
        ) : (
          <CustomButton
            onPress={() => JoinEvent(Number(id))}
            title="Я иду!"
            type={'action'}
            fill={'solid'}
          />
        )}
      </View>

      <StatusBar style={Platform.OS === 'ios' ? 'light' : 'auto'} />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    height: '100%',
    display: 'flex',
    width: '100%',
    marginBottom: 32,
  },
  close: {
    position: 'absolute',
    top: 16,
    right: 16,
    zIndex: 5,
  },
  content: {
    backgroundColor: '#ffffff',
    height: 555,
    width: '100%',
    paddingHorizontal: 24,
    paddingVertical: 16,
  },
  tags: {
    display: 'flex',
    flexDirection: 'row',
    gap: 8,
  },
  join: {
    borderRadius: 6,
    paddingVertical: 6,
    paddingHorizontal: 8,
    borderColor: colors.success,
    borderWidth: 1,
    width: 'auto',
  },
  join__text: {
    fontFamily: fonts.Unbounded,
    color: colors.success,
    fontWeight: 600,
    fontSize: 8,
  },
  public: {
    borderRadius: 6,
    paddingVertical: 6,
    paddingHorizontal: 8,
    borderColor: colors.primary,
    borderWidth: 1,
  },
  public__text: {
    fontFamily: fonts.Unbounded,
    color: colors.primary,
    fontWeight: 600,
    fontSize: 8,
  },
  mainText: {
    height: 350,
  },
  title: {
    marginTop: 8,
    color: colors.black,
    fontFamily: fonts.Unbounded,
    fontWeight: 800,
    fontSize: 20,
  },
  category: {
    color: colors.secondary,
    fontFamily: fonts.Montserrat,
    fontWeight: 600,
    fontSize: 12,
    marginTop: 2,
  },
  description: {
    color: colors.secondary,
    fontFamily: fonts.Montserrat,
    fontWeight: 400,
    fontSize: 12,
    marginTop: 16,
  },
  date: {
    display: 'flex',
    flexDirection: 'row',
    gap: 8,
  },
  date__item: {
    borderRadius: 6,
    backgroundColor: colors.primary,
    display: 'flex',
    justifyContent: 'center',
    alignContent: 'center',
    paddingHorizontal: 8,
    paddingVertical: 6,
  },
  item__text: {
    fontFamily: fonts.Unbounded,
    fontWeight: 600,
    fontSize: 10,
    color: '#ffffff',
  },
  location: {
    marginTop: 8,
    display: 'flex',
    flexDirection: 'row',
    gap: 8,
    alignItems: 'flex-start',
  },
  location__text: {
    fontFamily: fonts.Montserrat,
    fontWeight: 500,
    fontSize: 12,
    color: colors.secondary,
  },
  button: {
    bottom: 24,
    paddingHorizontal: 24,
  },
})

export default ModalScreen
