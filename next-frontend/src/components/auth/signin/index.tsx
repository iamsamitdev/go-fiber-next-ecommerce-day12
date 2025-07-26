"use client"

import Link from "next/link"
import React, { useState } from "react"
import { useAuth } from "@/hooks/useAuth"

const Signin = () => {
  const { loginForm, login, isLoading } = useAuth()
  const [message, setMessage] = useState<{ type: 'success' | 'error', text: string } | null>(null)

  const onSubmit = async (data: any) => {
    setMessage(null)
    const result = await login(data)
    
    if (result.success) {
      setMessage({ type: 'success', text: result.message })
    } else {
      setMessage({ type: 'error', text: result.message })
    }
  }

  return (
    <>
      <section className="flex items-center justify-center">
        <div className="w-full max-w-md">
          <div className="bg-white rounded-lg shadow-lg p-8">
            <div className="text-center mb-8">
              <h2 className="font-semibold text-2xl text-dark mb-2">
                เข้าสู่ระบบ
              </h2>
              <p className="text-gray-600">กรอกข้อมูลด้านล่าง</p>
            </div>

            {message && (
              <div className={`mb-5 p-4 rounded-lg ${
                message.type === 'success' 
                  ? 'bg-green-50 text-green-700 border border-green-200' 
                  : 'bg-red-50 text-red-700 border border-red-200'
              }`}>
                {message.text}
              </div>
            )}

            <div>
              <form onSubmit={loginForm.handleSubmit(onSubmit)}>
                <div className="mb-5">
                  <label htmlFor="email" className="block mb-2.5">
                    อีเมล <span className="text-red">*</span>
                  </label>

                  <input
                    {...loginForm.register('email')}
                    type="text"
                    id="email"
                    placeholder="กรอกอีเมล"
                    className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                  />
                  {loginForm.formState.errors.email && (
                    <p className="text-red-500 text-sm mt-1">
                      {loginForm.formState.errors.email.message}
                    </p>
                  )}
                </div>

                <div className="mb-5">
                  <label htmlFor="password" className="block mb-2.5">
                    รหัสผ่าน <span className="text-red">*</span>
                  </label>

                  <input
                    {...loginForm.register('password')}
                    type="password"
                    id="password"
                    placeholder="กรอกรหัสผ่าน"
                    autoComplete="current-password"
                    className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                  />
                  {loginForm.formState.errors.password && (
                    <p className="text-red-500 text-sm mt-1">
                      {loginForm.formState.errors.password.message}
                    </p>
                  )}
                </div>

                <button
                  type="submit"
                  disabled={isLoading}
                  className="w-full flex justify-center font-medium text-white bg-dark py-3 px-6 rounded-lg ease-out duration-200 hover:bg-blue mt-7.5 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {isLoading ? 'กำลังเข้าสู่ระบบ...' : 'เข้าสู่ระบบ'}
                </button>

                <Link
                  href="/auth/forgotpassword"
                  className="block text-center text-dark-4 mt-4.5 ease-out duration-200 hover:text-dark"
                >
                  ลืมรหัสผ่าน?
                </Link>

                <p className="text-center mt-6">
                  ยังไม่มีบัญชี?
                  <Link
                    href="/auth/signup"
                    className="text-dark ease-out duration-200 hover:text-blue pl-2"
                  >
                    สมัครสมาชิกทันที!
                  </Link>
                </p>
              </form>
            </div>
          </div>
        </div>
      </section>
    </>
  )
}

export default Signin
