import { useFonts } from "expo-font";

const useAppFonts = () => {
  const [fontsLoaded] = useFonts({
    Unbounded: require("../assets/fonts/Unbounded.ttf"),
    Montserrat: require("../assets/fonts/Montserrat.ttf"),
  });

  return fontsLoaded;
};

export default useAppFonts;
