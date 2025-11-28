"use server"

import { cookies } from "next/headers"

export async function signupAction(prevState: any, formData: FormData) {
 const givenName = formData.get('given-name') as string
 const familyName = formData.get('family-name') as string
 const email = formData.get('email') as string
 const password = formData.get('password') as string
 const confirmPassword = formData.get('confirm-password') as string

 if (!givenName || !familyName || !email || !password || !confirmPassword) {
  return { error: 'All fields are required.' }
 }

 if (password !== confirmPassword) {
  return { error: 'Passwords do not match.' }
 }

 try {
  const response = await fetch(new URL(process.env.AUTH_SERVICE_SIGNUP_ENDPOINT || ''), {
   method: 'POST',
   headers: { 'Content-Type': 'application/json' },
   body: JSON.stringify({ given_name: givenName, family_name: familyName, email, password, confirm_password: confirmPassword }),
  })

  if (!response.ok) {
   const error = await response.json()
   return { error: error.error || 'Signup failed' }
  }

  return { success: true, data: { email } }

 } catch (error) {
  return { error: 'Network error. Please try again.' }
 }
}

export async function verifyOTPAction(prevState: any, formData: FormData) {
 const otpCode = formData.get('otp') as string
 const email = formData.get('email') as string

 if (!otpCode) {
  return { error: 'OTP code is required.' }
 }

 try {
  const response = await fetch(new URL(process.env.AUTH_SERVICE_OTP_VALIDATE_ENDPOINT || ''), {
   method: 'POST',
   headers: { 'Content-Type': 'application/json' },
   body: JSON.stringify({ email, code: otpCode }),
  })

  if (!response.ok) {
   const error = await response.json()
   return { error: error.error || 'OTP verification failed' }
  }

  return { success: true }

 } catch (error) {
  return { error: 'It was not possible to verify the OTP code. Please try again or click on resend code.' }
 }
}

export async function signinAction(prevState: any, formData: FormData) {
 const email = formData.get('email') as string
 const password = formData.get('password') as string

 if (!email || !password) {
  return { error: 'Email and password are required.' }
 }

 try {
  const response = await fetch(new URL(process.env.AUTH_SERVICE_LOGIN_ENDPOINT || ''), {
   method: 'POST',
   headers: { 'Content-Type': 'application/json' },
   body: JSON.stringify({ email, password }),
  })

  if (!response.ok) {
   const error = await response.json()
   return { error: error.error || 'Login failed' }
  }

  const data = await response.json()
  const cookieStore = await cookies()
  cookieStore.set('session_token', data.token, { httpOnly: true, path: '/' })

  return { success: true, data }
 } catch (error) {
  return { error: 'Network error. Please try again.' }
 }
}

export async function signoutAction() {
 const cookieStore = await cookies()
 const token = cookieStore.get('session_token')?.value

 if (!token) {
  return { error: 'No active session found.' }
 }

 try {
  const response = await fetch(new URL(process.env.AUTH_SERVICE_SIGNOUT_ENDPOINT || ''), {
   method: 'POST',
   headers: { 'Content-Type': 'application/json', 'Authorization': `${token}` },
  })

  if (!response.ok) {
   const error = await response.json()
   return { error: error.error || 'Signout failed' }
  }

  cookieStore.delete('session_token')

  return { success: true }
 } catch (error) {
  return { error: 'Network error. Please try again.' }
 }
}