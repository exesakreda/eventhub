import { StyleSheet, TouchableOpacity, View } from 'react-native'
import { Text } from 'react-native'
import { colors } from '@/constants/colors'
import * as Haptics from 'expo-haptics'
import { fonts } from '@/constants/fonts'

const handlePress = () => {
  Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light)
}

const CustomButton = ({
  onPress,
  title,
  type,
  fill,
}: {
  onPress: any
  title: string
  type: 'action' | 'cancel'
  fill: 'solid' | 'bordered'
}) => {
  const styles = StyleSheet.create({
    container: {
      width: '100%',
    },
    button_solid: {
      boxSizing: 'border-box',
      backgroundColor: type == 'action' ? colors.primary : colors.error,
      width: '100%',
      height: 48,
      borderRadius: 12,
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
    },
    text_solid: {
      color: '#ffffff',
      fontSize: 12,
      fontWeight: 600,
      fontFamily: fonts.Unbounded,
    },
    button_bordered: {
      boxSizing: 'border-box',
      backgroundColor: '#ffffff',
      borderColor: type == 'action' ? colors.primary : colors.error,
      borderWidth: 2,
      width: '100%',
      height: 48,
      borderRadius: 12,
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
    },
    text_bordered: {
      color: type == 'action' ? colors.primary : colors.error,
      fontSize: 12,
      fontWeight: 600,
      fontFamily: fonts.Unbounded,
    },
  })
  return (
    <View style={styles.container}>
      <TouchableOpacity
        onPress={() => {
          onPress()
          handlePress()
        }}
        style={fill == 'solid' ? styles.button_solid : styles.button_bordered}
      >
        <Text
          style={fill == 'solid' ? styles.text_solid : styles.text_bordered}
        >
          {title}
        </Text>
      </TouchableOpacity>
    </View>
  )
}

export default CustomButton
