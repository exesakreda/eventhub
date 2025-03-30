import { AuthContext } from '@/app/(auth)/AuthContext'
import AsyncStorage from '@react-native-async-storage/async-storage'
import { router } from 'expo-router'
import { useCallback, useContext, useEffect, useState } from 'react'
import { Button, Text } from 'react-native'
import { SafeAreaView } from 'react-native-safe-area-context'

const Profile = () => {
  const auth = useContext(AuthContext)
  const [firstName, setFirstName] = useState('no first name')
  const [lastName, setLastName] = useState('no last name')
  const [username, setUsername] = useState('no username')

  const getUserData = useCallback(async () => {
    try {
      const token = await AsyncStorage.getItem('token')
      if (!token) throw new Error('Отсутствует токен')

      const response = await fetch('http://192.168.0.106:8080/getuserdata', {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      })

      if (!response.ok) {
        if (response.status === 401) {
          router.replace('/(auth)/login')
        }
        throw new Error(`Ошибка запроса: ${response.status}`)
      }

      const data = await response.json()
      setFirstName(data.first_name || 'no first name')
      setLastName(data.last_name || 'no last name')
      setUsername(data.username || 'no username')
    } catch (error) {
      console.error('Ошибка запроса:', error)
    }
  }, [auth])

  useEffect(() => {
    getUserData()
  }, [getUserData])
  return (
    <SafeAreaView>
      <Text>first name: {firstName}</Text>
      <Text>last name: {lastName}</Text>
      <Text>username: {username}</Text>

      <Button
        title="Logout"
        onPress={() => {
          AsyncStorage.removeItem('token')
          router.replace('/(auth)/login')
        }}
      />
    </SafeAreaView>
  )
}

export default Profile
