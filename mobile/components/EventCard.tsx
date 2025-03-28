import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { Image, StyleSheet, Text, View } from "react-native";

type EventCardProps = {
  date: string;
  start_time: string;
  end_time: string;
  title: string;
  category: string;
  place: string;
};

const EventCard = ({
  date,
  start_time,
  end_time,
  title,
  category,
  place,
}: EventCardProps) => {
  return (
    <View style={styles.card}>
      <Image
        source={require("../assets/images/blank_image_card.png")}
        style={{ width: "100%", height: 120 }}
        resizeMode="contain"
      />
      <View style={styles.card__content}>
        <View style={styles.content__date}>
          <View style={styles.date__day}>
            <Text style={styles.date__text}>{date}</Text>
          </View>
          <View style={styles.date__time}>
            <Text style={styles.date__text}>
              {start_time} - {end_time}
            </Text>
          </View>
        </View>
        <Text style={styles.content__title}>{title}</Text>
        <Text style={styles.content__category}>{category}</Text>
        <View style={styles.content__place}>
          <Image
            source={require("../assets/icons/location.png")}
            style={{ width: 10, height: 13 }}
            resizeMode="contain"
          />
          <Text style={styles.place__text}>{place}</Text>
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  card: {
    width: 250,
    height: 282,
    borderRadius: 16,
    backgroundColor: "#F8F9FE",
    overflow: "hidden",
    boxShadow:
      "rgba(14, 63, 126, 0.03) 0px 0px 0px 1px, rgba(42, 51, 69, 0.03) 0px 1px 1px -0.5px, rgba(42, 51, 70, 0.03) 0px 3px 3px -1.5px, rgba(42, 51, 70, 0.03) 0px 6px 6px -3px, rgba(14, 63, 126, 0.03) 0px 12px 12px -6px, rgba(14, 63, 126, 0.03) 0px 24px 24px -12px",
    marginBottom: 32,
  },
  card__content: {
    padding: 16,
    boxSizing: "border-box",
    flexGrow: 1,
  },
  content__date: {
    display: "flex",
    flexDirection: "row",
  },
  date__day: {
    paddingHorizontal: 8,
    paddingVertical: 6,
    borderStyle: "solid",
    borderColor: colors.primary,
    borderWidth: 1,
    borderRadius: 6,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  date__time: {
    paddingHorizontal: 8,
    paddingVertical: 6,
    borderStyle: "solid",
    borderColor: colors.primary,
    borderWidth: 1,
    borderRadius: 6,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    marginLeft: 8,
  },
  date__text: {
    fontWeight: 600,
    fontFamily: fonts.Unbounded,
    fontSize: 8,
    textTransform: "uppercase",
    color: colors.primary,
  },
  content__title: {
    fontWeight: 700,
    fontFamily: fonts.Unbounded,
    marginTop: 8,
    color: colors.black,
    fontSize: 14,
  },
  content__category: {
    fontWeight: 600,
    fontFamily: fonts.Montserrat,
    fontSize: 12,
    color: colors.secondary,
    marginTop: 4,
  },
  content__place: {
    marginTop: "auto",
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
  },
  place__text: {
    marginLeft: 8,
    fontFamily: fonts.Montserrat,
    fontSize: 12,
    color: colors.secondary,
    fontWeight: 500,
  },
});

export default EventCard;
