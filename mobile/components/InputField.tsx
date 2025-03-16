import { colors } from "@/constants/colors";
import { useState } from "react";
import {
  StyleSheet,
  TextInput,
  View,
  Image,
  TouchableOpacity,
} from "react-native";
import * as Haptics from "expo-haptics";
import { fonts } from "@/constants/fonts";

type InputFieldProps = {
  value: string;
  onChangeText: (text: string) => void;
  placeholder: string;
  isIncorrect?: boolean;
  secureTextEntry?: boolean;
  props?: any;
};

const InputField = ({
  value,
  onChangeText,
  placeholder,
  isIncorrect,
  secureTextEntry,
  ...props
}: InputFieldProps) => {
  const [isPasswordVisible, setIsPasswordVisible] = useState(secureTextEntry);

  return (
    <View style={styles.container}>
      <TextInput
        style={[styles.input, isIncorrect && { borderColor: colors.error }]}
        value={value}
        onChangeText={onChangeText}
        placeholder={placeholder}
        placeholderTextColor={colors.form_text}
        secureTextEntry={isPasswordVisible}
        {...props}
      ></TextInput>
      {secureTextEntry ? (
        <TouchableOpacity
          onPress={() => {
            setIsPasswordVisible(!isPasswordVisible);
            Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Soft);
          }}
          style={styles.passwordVisibility}
        >
          <Image
            source={
              isPasswordVisible
                ? require("../assets/icons/password_hide.png")
                : require("../assets/icons/password_show.png")
            }
            style={{ width: 16, height: 16 }}
          />
        </TouchableOpacity>
      ) : (
        <View></View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    width: "100%",
  },
  input: {
    width: "100%",
    height: 48,
    color: colors.form_text,
    borderColor: colors.form_border,
    borderWidth: 1,
    borderRadius: 12,
    paddingHorizontal: 16,
    paddingVertical: 14,
    boxSizing: "border-box",
    fontFamily: fonts.Montserrat,
    fontSize: 14,
  },
  passwordVisibility: {
    position: "absolute",
    marginTop: 16,
    right: 16,
  },
});

export default InputField;
