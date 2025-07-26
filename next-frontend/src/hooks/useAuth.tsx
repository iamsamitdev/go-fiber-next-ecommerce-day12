"use client"

import { useState, useEffect } from "react"
import { useRouter } from "next/navigation"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import axios from "axios"
import Cookies from "js-cookie"

// Zod schema for form validation
const registerSchema = z
  .object({
    firstName: z.string().min(2, "ชื่อต้องมีอย่างน้อย 2 ตัวอักษร"),
    lastName: z.string().min(2, "นามสกุลต้องมีอย่างน้อย 2 ตัวอักษร"),
    email: z.string().email("รูปแบบอีเมลไม่ถูกต้อง"),
    password: z.string().min(6, "รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร"),
    confirmPassword: z.string(),
    phone: z.string().min(10, "เบอร์โทรศัพท์ต้องมีอย่างน้อย 10 หลัก"),
    address: z.string().min(10, "ที่อยู่ต้องมีอย่างน้อย 10 ตัวอักษร"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "รหัสผ่านไม่ตรงกัน",
    path: ["confirmPassword"],
  })

// Zod schema for login validation
const loginSchema = z.object({
  email: z.string().email("รูปแบบอีเมลไม่ถูกต้อง"),
  password: z.string().min(1, "กรุณากรอกรหัสผ่าน"),
})

// Zod schema for forgot password validation
const forgotPasswordSchema = z.object({
  email: z.string().email("รูปแบบอีเมลไม่ถูกต้อง"),
})

// Types
export type RegisterFormData = z.infer<typeof registerSchema>
export type LoginFormData = z.infer<typeof loginSchema>
export type ForgotPasswordFormData = z.infer<typeof forgotPasswordSchema>

// Interface User
interface User {
  id: string
  first_name: string
  last_name: string
  email: string
  phone: string
  address: string
  avatar: string
  active: boolean
  role_id: string
  role: {
    id: string
    name: string
    description: string
    created_at: string
    updated_at: string
  }
  created_at: string
  updated_at: string
}

// Response interface
interface AuthResponse {
  success: boolean
  message: string
  data?: {
    user: User
    token: string
    refresh_token: string
  }
}

// Axios instance
const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
})

// use Interceptors to handle token
api.interceptors.request.use((config) => {
  const token = Cookies.get("authToken")
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// สร้าง hook สำหรับการจัดการการยืนยันตัวตน
export const useAuth = () => {
  const [user, setUser] = useState<User | null>(null) // สถานะผู้ใช้
  const [isLoading, setIsLoading] = useState(false) // สถานะการโหลด
  const [isAuthenticated, setIsAuthenticated] = useState(false) // สถานะการยืนยันตัวตน
  const router = useRouter() // ใช้สำหรับการนำทาง

  // check if user is authenticated on mount
  useEffect(() => {
    const token = Cookies.get("authToken")
    const userData = Cookies.get("userData")

    if (token && userData) {
      try {
        const parsedUser = JSON.parse(userData)
        setUser(parsedUser)
        setIsAuthenticated(true)
      } catch (error) {
        console.error("Error parsing user data:", error)
        logout()
      }
    }
  }, [])

  // Register form
  const registerForm = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      confirmPassword: "",
      phone: "",
      address: "",
    },
  })

  // Login form
  const loginForm = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  })

  // Forgot password form
  const forgotPasswordForm = useForm<ForgotPasswordFormData>({
    resolver: zodResolver(forgotPasswordSchema),
    defaultValues: {
      email: "",
    },
  })

  // Function to register a new user
  const register = async (data: RegisterFormData) => {
    setIsLoading(true)
    try {
      const response = await api.post<AuthResponse>("/api/v1/auth/register", {
        first_name: data.firstName,
        last_name: data.lastName,
        email: data.email,
        password: data.password,
        phone: data.phone,
        address: data.address,
      })

      if (response.data.success && response.data.data) {
        // Redirect to signin page
        router.push("/auth/signin")

        return { success: true, message: response.data.message }
      }

      return {
        success: false,
        message: response.data.message || "เกิดข้อผิดพลาดในการสมัครสมาชิก",
      }
    } catch (error: any) {
      console.error("Register error:", error)
      const errorMessage =
        error.response?.data?.message || "เกิดข้อผิดพลาดในการสมัครสมาชิก"
      return { success: false, message: errorMessage }
    } finally {
      setIsLoading(false)
    }
  }

  // Function to login a user
  const login = async (data: LoginFormData) => {
    setIsLoading(true)
    try {
      const response = await api.post<AuthResponse>("/api/v1/auth/login", {
        email: data.email,
        password: data.password,
      })

      if (response.data.success && response.data.data) {
        // Save token and user data to cookies (7 days from now)
        Cookies.set("authToken", response.data.data.token, {
          expires: 7, // Simple 7 days
          secure: process.env.NODE_ENV === "production",
          sameSite: "strict",
        })
        Cookies.set('refreshToken', response.data.data.refresh_token, {
          expires: 30, // 30 days
          secure: process.env.NODE_ENV === 'production',
          sameSite: 'strict'
        })
        Cookies.set("userData", JSON.stringify(response.data.data.user), {
          expires: 7, // Simple 7 days  
          secure: process.env.NODE_ENV === "production",
          sameSite: "strict",
        })

        setUser(response.data.data.user)
        setIsAuthenticated(true)
        loginForm.reset()

        // Redirect base on user role
        const userRole = response.data.data.user.role.name
        if (userRole === "admin") {
          router.push("/admin/dashboard")
        } else {
          router.push("/") // หน้าแรกสำหรับ user ทั่วไป
        }

        return { success: true, message: response.data.message }
      }

      return {
        success: false,
        message: response.data.message || "เกิดข้อผิดพลาดในการเข้าสู่ระบบ",
      }
    } catch (error: any) {
      console.error("Login error:", error)
      const errorMessage =
        error.response?.data?.message || "เกิดข้อผิดพลาดในการเข้าสู่ระบบ"
      return { success: false, message: errorMessage }
    } finally {
      setIsLoading(false)
    }
  }

  // Forgot password function
  const forgotPassword = async (data: ForgotPasswordFormData) => {
    setIsLoading(true)
    try {
      const response = await api.post<AuthResponse>('/api/v1/auth/forgot-password', {
        email: data.email
      })

      if (response.data.success) {
        forgotPasswordForm.reset()
        return { success: true, message: response.data.message || 'ส่งลิงก์รีเซ็ตรหัสผ่านไปยังอีเมลของคุณแล้ว' }
      }

      return { success: false, message: response.data.message || 'เกิดข้อผิดพลาดในการส่งลิงก์รีเซ็ตรหัสผ่าน' }
    } catch (error: any) {
      console.error('Forgot password error:', error)
      const errorMessage = error.response?.data?.message || 'เกิดข้อผิดพลาดในการส่งลิงก์รีเซ็ตรหัสผ่าน'
      return { success: false, message: errorMessage }
    } finally {
      setIsLoading(false)
    }
  }

  // Logout function
  const logout = () => {
    Cookies.remove('authToken')
    Cookies.remove('userData')
    setUser(null)
    setIsAuthenticated(false)
    router.push('/auth/signin')
  }

  // Helper functions
  const isAdmin = () => {
    return user?.role?.name === 'admin'
  }

  const isUser = () => {
    return user?.role?.name === 'user'
  }

  const getUserFullName = () => {
    if (!user) return ''
    return `${user.first_name} ${user.last_name}`.trim()
  }

  return {
    // State
    user,
    isLoading,
    isAuthenticated,
    
    // Forms
    registerForm,
    loginForm,
    forgotPasswordForm,
    
    // Functions
    register,
    login,
    forgotPassword,
    logout,

    // Helper functions
    isAdmin,
    isUser,
    getUserFullName

  }

}
