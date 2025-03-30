import { Tabs } from "expo-router";
import { Image, Text, View } from "react-native";

const TabIcon = ({ source }: { source: any }) => (
  <Image
    source={source}
    style={{ width: 20, height: 20 }}
    resizeMode="contain"
  />
);

const icons = {
  home: require("../../../assets/icons/home.png"),
  home_focused: require("../../../assets/icons/home_focused.png"),
  my_events: require("../../../assets/icons/my_events.png"),
  my_events_focused: require("../../../assets/icons/my_events_focused.png"),
  notifications: require("../../../assets/icons/notifications.png"),
  notifications_focused: require("../../../assets/icons/notifications_focused.png"),
  profile: require("../../../assets/icons/profile.png"),
  profile_focused: require("../../../assets/icons/profile_focused.png"),
};

const Layout = () => (
  <Tabs
    initialRouteName="index"
    screenOptions={{
      tabBarActiveTintColor: "white",
      tabBarLabelStyle: { display: "none" },
      tabBarStyle: {
        height: 78,
        paddingTop: 8,
      },
    }}
  >
    <Tabs.Screen
      name="home"
      options={{
        title: "Главная",
        headerShown: false,
        tabBarIcon: ({ focused }) => (
          <TabIcon source={focused ? icons.home_focused : icons.home} />
        ),
      }}
    />
    <Tabs.Screen
      name="my_events"
      options={{
        title: "Мои ивенты",
        headerShown: false,
        tabBarIcon: ({ focused }) => (
          <TabIcon
            source={focused ? icons.my_events_focused : icons.my_events}
          />
        ),
      }}
    />
    <Tabs.Screen
      name="notifications"
      options={{
        title: "Уведомления",
        headerShown: false,
        tabBarIcon: ({ focused }) => (
          <TabIcon
            source={focused ? icons.notifications_focused : icons.notifications}
          />
        ),
      }}
    />
    <Tabs.Screen
      name="profile"
      options={{
        title: "Профиль",
        headerShown: false,
        tabBarIcon: ({ focused }) => (
          <TabIcon source={focused ? icons.profile_focused : icons.profile} />
        ),
      }}
    />
  </Tabs>
);

export default Layout;
