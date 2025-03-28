import { Stack } from "expo-router";

const Layout = () => {
  return (
    <Stack>
      <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
      <Stack.Screen
        name="eventModal"
        options={{
          presentation: "modal",
          headerShown: false,
          gestureEnabled: true,
          // modalPresentationStyle: "pageSheet",
        }}
      />
    </Stack>
  );
};

export default Layout;
