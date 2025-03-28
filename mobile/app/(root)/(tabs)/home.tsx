import { AuthContext } from "@/app/(auth)/AuthContext";
import EventCard from "@/components/EventCard";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { Link, router } from "expo-router";
import { useCallback, useContext, useEffect, useState } from "react";
import {
  Button,
  Dimensions,
  Image,
  Modal,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { format, parseISO } from "date-fns";
import { ru, tr } from "date-fns/locale";

type Event = {
  id: number;
  title: string;
  description: string;
  category: string;
  is_public: boolean;
  status: string;
  date: string;
  start_time: string;
  end_time: string;
  location: string;
  creator_id: number;
  organization_id: number;
};

export default function Home() {
  const auth = useContext(AuthContext);

  const [activeJoinedEvents, setActiveJoinedEvents] = useState<Event[]>([]);
  const [activeEvents, setActiveEvents] = useState<Event[]>([]);

  const [loadingActiveJoined, setLoadingActiveJoined] = useState(false);
  const [loadingActive, setLoadingActive] = useState(false);

  const formatDate = (date: string) => {
    const parsedDate = parseISO(date);
    const formattedDate = format(parsedDate, "d MMMM", {
      locale: ru,
    }).toUpperCase();

    return formattedDate;
  };

  const formatTime = (time: string) => {
    if (!time) return "00:00";
    return time.slice(0, 5);
  };

  const LoadActiveJoinedEvents = useCallback(async () => {
    const timeout = setTimeout(() => setLoadingActiveJoined(true), 300);
    try {
      const token = await AsyncStorage.getItem("token");
      if (!token) throw new Error("Отсутствует токен");

      const response = await fetch(
        "http://192.168.0.106:8080/getevents?status=active&role=participant",
        {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        },
      );

      if (!response.ok) {
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        if (response.status === 404) {
          setActiveJoinedEvents([]);
          return;
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }

      const data = await response.json();
      setActiveJoinedEvents(data);
    } catch (error) {
      console.error("Ошибка запроса:", error);
    } finally {
      clearTimeout(timeout);
      setLoadingActiveJoined(false);
    }
  }, [auth]);

  const LoadAllActiveEvents = useCallback(async () => {
    const timeout = setTimeout(() => setLoadingActive(true), 300);
    try {
      const token = await AsyncStorage.getItem("token");
      if (!token) throw new Error("Отсутствует токен");

      const response = await fetch(
        "http://192.168.0.106:8080/getevents?status=active",
        {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        },
      );

      if (!response.ok) {
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        if (response.status === 404) {
          setActiveEvents([]);
          return;
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }

      const data = await response.json();
      setActiveEvents(data);
    } catch (error) {
      console.error("Ошибка запроса:", error);
    } finally {
      clearTimeout(timeout);
      setLoadingActive(false);
    }
  }, [auth]);

  useEffect(() => {
    LoadActiveJoinedEvents();
    LoadAllActiveEvents();
    const interval = setInterval(() => {
      LoadActiveJoinedEvents();
      LoadAllActiveEvents();
    }, 30000);

    return () => clearInterval(interval);
  }, [LoadActiveJoinedEvents, LoadAllActiveEvents]);

  const renderActiveJoinedEvents = () => {
    return activeJoinedEvents.map((event) => {
      return (
        <View style={styles.category__item} key={event.id}>
          <EventCard
            date={formatDate(event.date)}
            start_time={formatTime(event.start_time)}
            end_time={formatTime(event.end_time)}
            title={event.title}
            category={event.category}
            place={event.location}
          />
        </View>
      );
    });
  };

  const renderActiveEvents = () => {
    return activeEvents.map((event) => {
      return (
        <View style={styles.category__item} key={event.id}>
          <EventCard
            date={formatDate(event.date)}
            start_time={formatTime(event.start_time)}
            end_time={formatTime(event.end_time)}
            title={event.title}
            category={event.category}
            place={event.location}
          />
        </View>
      );
    });
  };

  const screenHeight = Dimensions.get("window").height;

  return (
    <SafeAreaView style={{ backgroundColor: "#ffffff" }}>
      <View style={styles.header}>
        <View style={styles.header__content}>
          <Image
            source={require("../../../assets/icons/logo.png")}
            style={{ width: 20, height: 20 }}
            resizeMode="contain"
          />
          <Text style={styles.header__title}>EVENTHUB</Text>
        </View>
      </View>

      <Button title="Открыть PageSheet" onPress={() => router.push("/modal")} />

      <ScrollView style={styles.categories}>
        {activeJoinedEvents.length > 0 ? (
          <View style={styles.category}>
            <Text style={styles.category__title}>Вы участвуете</Text>
            <ScrollView
              horizontal={true}
              showsHorizontalScrollIndicator={false}
              style={styles.category__content}
            >
              {renderActiveJoinedEvents()}
            </ScrollView>
          </View>
        ) : (
          <View></View>
        )}

        {activeEvents.length > 0 ? (
          <View style={styles.category}>
            <Text style={styles.category__title}>Все мероприятия</Text>
            <ScrollView
              horizontal={true}
              showsHorizontalScrollIndicator={false}
              style={styles.category__content}
            >
              {renderActiveEvents()}
            </ScrollView>
          </View>
        ) : (
          <View></View>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  header: {
    position: "fixed",
    height: 46,
    display: "flex",
    alignItems: "center",
    paddingTop: 4,
    boxSizing: "border-box",
  },
  header__content: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "center",
    alignContent: "center",
  },
  header__title: {
    fontFamily: fonts.Unbounded,
    fontWeight: 800,
    color: colors.primary,
    marginLeft: 8,
    fontSize: 16,
  },
  categories: {
    width: "100%",
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
});
