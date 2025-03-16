import InputField from "@/components/InputField";
import SolidButton from "@/components/SolidButton";
import { colors } from "@/constants/colors";
import { Link, router } from "expo-router";
import { useContext, useState } from "react";
import {
  Keyboard,
  KeyboardAvoidingView,
  Platform,
  StyleSheet,
  Text,
  TouchableWithoutFeedback,
} from "react-native";
import { View, ScrollView } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { AuthContext } from "./AuthContext";
import { fonts } from "@/constants/fonts";

const Login = () => {
  const auth = useContext(AuthContext);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async () => {
    try {
      const response = await fetch("http://192.168.0.105:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });
      const data = await response.json();
      if (response.ok && auth) {
        auth.login(data.token);
        router.replace("/(root)/(tabs)/home");
      } else {
        setError("Неверные данные");
      }
    } catch (error) {
      setError("Ошибка сервера");
    }
  };

  return (
    <SafeAreaView style={styles.content}>
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "position"}
        style={{ flex: 1 }}
      >
        <TouchableWithoutFeedback onPress={() => Keyboard.dismiss()}>
          <ScrollView
            contentContainerStyle={styles.scrollView}
            keyboardShouldPersistTaps="handled"
            scrollEnabled={false}
          >
            <Text style={styles.title}>Добро пожаловать! 👋</Text>
            <View style={styles.form}>
              <View style={styles.form__element}>
                <InputField
                  value={username}
                  placeholder="Имя пользователя"
                  isIncorrect={error !== ""}
                  secureTextEntry={false}
                  onChangeText={(text) => setUsername(text)}
                />
              </View>
              <View style={styles.form__element}>
                <InputField
                  value={password}
                  placeholder="Пароль"
                  isIncorrect={error !== ""}
                  secureTextEntry={true}
                  onChangeText={(text) => setPassword(text)}
                />
              </View>
            </View>
            <SolidButton onPress={handleLogin} title="Войти" />
            <View style={styles.registration}>
              <Text style={styles.registration__title}>
                Нет аккаунта?{" "}
                <Link
                  href="/(auth)/registration"
                  style={styles.registration__link}
                >
                  <Text>Зарегистрироваться</Text>
                </Link>
              </Text>
            </View>
          </ScrollView>
        </TouchableWithoutFeedback>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  content: {
    flex: 1,
    width: "100%",
    paddingHorizontal: 24,
  },
  scrollView: {
    flexGrow: 1,
    justifyContent: "flex-end",
  },
  title: {
    fontSize: 24,
    fontWeight: 800,
    color: colors.black,
    fontFamily: fonts.Unbounded,
  },
  form: {
    marginTop: 24,
  },
  registration: {
    marginTop: 16,
    width: "100%",
    display: "flex",
    alignItems: "center",
  },
  registration__title: {
    fontSize: 12,
    color: colors.secondary,
    fontWeight: 400,
    fontFamily: fonts.Montserrat,
  },
  registration__link: {
    color: colors.primary,
    fontWeight: 600,
  },
  form__element: {
    marginBottom: 16,
  },
});

export default Login;
