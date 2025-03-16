import EventCard from "@/components/EventCard";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { Link } from "expo-router";
import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Home() {
  return (
    <SafeAreaView>
      <View style={styles.header}>
        <View style={styles.header__content}>
          <Image
            source={require("../../../assets/icons/logo.png")}
            style={{ width: 20, height: 20 }}
            resizeMode="contain"
          />
          <Text style={styles.header__title}>EVENTHUB</Text>s
        </View>
      </View>
      <ScrollView style={styles.categories}>
        <View style={styles.category}>
          <Text style={styles.category__title}>Вы участвуете</Text>
          <ScrollView
            horizontal={true}
            showsHorizontalScrollIndicator={false}
            style={styles.category__content}
          >
            <View style={styles.catergory__item}>
              <EventCard
                date="16 марта"
                start_time="12:00"
                end_time="15:00"
                title="Граффити Москвы и Чебоксар для новых работников"
                category="Фестиваль уличного искусства"
                place="Винзавод, Москва"
              />
            </View>
            <View style={styles.catergory__item}>
              <EventCard
                date="16 марта"
                start_time="12:00"
                end_time="15:00"
                title="Граффити Москвы"
                category="Фестиваль уличного искусства"
                place="Винзавод, Москва"
              />
            </View>
            <View style={styles.catergory__item}>
              <EventCard
                date="16 марта"
                start_time="12:00"
                end_time="15:00"
                title="Граффити Москвы и Чебоксар для новых работников"
                category="Фестиваль уличного искусства"
                place="Винзавод, Москва"
              />
            </View>
          </ScrollView>
        </View>
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
    // height: ,
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
  catergory__item: {
    marginRight: 12,
  },
});
