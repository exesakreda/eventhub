import { Text, View, StyleSheet, Pressable } from 'react-native'
import { SafeAreaView } from 'react-native-safe-area-context'
import { useRef, useState } from 'react'
import Swiper from 'react-native-swiper'
import { onboarding } from '@/constants'
import { router } from 'expo-router'
import { colors } from '@/constants/colors'
import * as Haptics from 'expo-haptics'
import { fonts } from '@/constants/fonts'
import CustomButton from '@/components/CustomButton'

const Onboarding = () => {
  const swiperRef = useRef<Swiper>(null)
  const [activeIndex, setActiveIndex] = useState(0)
  const isLastSlide = activeIndex === onboarding.length - 1

  return (
    <SafeAreaView style={styles.content}>
      <SafeAreaView style={styles.skip}>
        <Pressable
          onPress={() => {
            router.replace('/(auth)/login')
            Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light)
          }}
          hitSlop={10}
        >
          <Text style={styles.skip__text}>Пропустить</Text>
        </Pressable>
      </SafeAreaView>
      <Swiper
        ref={swiperRef}
        // scrollEnabled={false}
        loop={false}
        onIndexChanged={(index) => {
          setActiveIndex(index)
        }}
      >
        {onboarding.map((item) => (
          <View style={styles.text_container} key={item.id}>
            <Text style={styles.title}>{item.title}</Text>
            <Text style={styles.description}>{item.description}</Text>
          </View>
        ))}
      </Swiper>
      <View style={styles.button}>
        <CustomButton
          title={isLastSlide ? 'Начать!' : 'Далее'}
          onPress={() => {
            isLastSlide
              ? router.replace('/(auth)/login')
              : swiperRef.current?.scrollBy(1)
          }}
          type={'action'}
          fill={'solid'}
        />
      </View>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  content: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  text_container: {
    height: 220,
    width: '100%',
    position: 'absolute',
    bottom: 0,
    paddingHorizontal: 24,
  },
  title: {
    fontSize: 24,
    fontWeight: '800',
    color: colors.black,
    fontFamily: fonts.Unbounded,
  },
  description: {
    fontSize: 12,
    fontWeight: '400',
    marginTop: 16,
    color: colors.secondary,
    fontFamily: fonts.Montserrat,
  },
  button: {
    width: '100%',
    paddingHorizontal: 24,
  },
  skip: {
    position: 'absolute',
    top: 8,
    right: 16,
    zIndex: 100,
  },
  skip__text: {
    color: colors.secondary,
    fontSize: 12,
  },
})

export default Onboarding
