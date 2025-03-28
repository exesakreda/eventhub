import { Redirect } from "expo-router";
import { useAuth } from "./(auth)/AuthContext";
import useAppFonts from "@/utils/useAppFonts";

const Home = () => {
  const { user } = useAuth();
  useAppFonts();

  // prod
  return <Redirect href={user ? "/(root)/(tabs)/home" : "/(auth)/welcome"} />;

  // dev
  // return <Redirect href={"/(root)/(tabs)/home"} />;
};

export default Home;
