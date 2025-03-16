import { StyleSheet, TouchableOpacity, View } from "react-native";
import { Text } from "react-native";
import { colors } from "@/constants/colors";
import * as Haptics from "expo-haptics";
import { fonts } from "@/constants/fonts";

const styles = StyleSheet.create({
  container: {
    width: "100%",
  },
  button: {
    boxSizing: "border-box",
    backgroundColor: colors.primary,
    width: "100%",
    height: 48,
    borderRadius: 12,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  text: {
    color: "#FFFFFF",
    fontSize: 12,
    fontWeight: 600,
    fontFamily: fonts.Unbounded,
  },
});

const handlePress = () => {
  Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Rigid);
};

const SolidButton = ({ onPress, title }: { onPress: any; title: string }) => (
  <View style={styles.container}>
    <TouchableOpacity
      onPress={() => {
        onPress();
        handlePress();
      }}
      style={styles.button}
    >
      <Text style={styles.text}>{title}</Text>
    </TouchableOpacity>
  </View>
);

export default SolidButton;
