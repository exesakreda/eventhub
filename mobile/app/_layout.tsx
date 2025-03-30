import { NavigationContainer } from "@react-navigation/native";
import { Slot, Stack } from "expo-router";
import * as SplashScreen from "expo-splash-screen";
import { useContext } from "react";
import { AuthProvider } from "./(auth)/AuthContext";
import { enableScreens } from "react-native-screens";

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  enableScreens();

  return (
    <AuthProvider>
      <Stack>
        <Stack.Screen name="index" options={{ headerShown: false }} />
        <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        <Stack.Screen name="(root)" options={{ headerShown: false }} />
        <Stack.Screen
          name="modal"
          options={{ presentation: "modal", headerShown: false }}
        />
      </Stack>
    </AuthProvider>
  );
}
