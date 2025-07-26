"use client"

import Link from "next/link"
import React, { useState } from "react"
import { useAuth } from "@/hooks/useAuth"

const Signup = () => {

  const  { registerForm, register, isLoading } = useAuth()
  const [message, setMessage] = useState<{ type: 'success' | 'error', text: string } | null>(null)

  const onSubmit = async (data: any) => {
    setMessage(null)
    const result = await register(data)
    
    if (result.success) {
      setMessage({ type: 'success', text: result.message })
    } else {
      setMessage({ type: 'error', text: result.message })
    }
  }

  return (
    <>
      <section className="flex items-center justify-center">
        <div className="w-full max-w-2xl">
          <div className="bg-white rounded-lg shadow-lg p-8">
            <div className="text-center mb-8">
              <h2 className="font-semibold text-2xl text-dark mb-2">
                สร้างบัญชีผู้ใช้
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


            <div className="mt-5.5">
              <form onSubmit={registerForm.handleSubmit(onSubmit)}>
                 {/* First Row - Name Fields */}
                 <div className="grid grid-cols-1 md:grid-cols-2 gap-5 mb-5">
                
                    <div>
                      <label htmlFor="firstName" className="block mb-2.5">
                        ชื่อ <span className="text-red">*</span>
                      </label>

                      <input
                        {...registerForm.register("firstName")}
                        type="text"
                        id="firstName"
                        placeholder="กรอกชื่อ"
                        className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                      />
                      {registerForm.formState.errors.firstName && (
                          <p className="text-red-500 text-sm mt-1">
                            {registerForm.formState.errors.firstName.message}
                          </p>
                      )}
                    </div>

                    <div>
                      <label htmlFor="lastName" className="block mb-2.5">
                        นามสกุล <span className="text-red">*</span>
                      </label>

                      <input
                        {...registerForm.register('lastName')}
                        type="text"
                        id="lastName"
                        placeholder="กรอกนามสกุล"
                        className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                      />
                      {registerForm.formState.errors.lastName && (
                        <p className="text-red-500 text-sm mt-1">
                          {registerForm.formState.errors.lastName.message}
                        </p>
                      )}
                    </div>

                </div>

                {/* Second Row - Email and Phone */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-5 mb-5">
                  <div>
                    <label htmlFor="email" className="block mb-2.5">
                      อีเมล <span className="text-red">*</span>
                    </label>

                    <input
                      {...registerForm.register('email')}
                      type="text"
                      id="email"
                      placeholder="กรอกอีเมล"
                      className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                    />
                    {registerForm.formState.errors.email && (
                      <p className="text-red-500 text-sm mt-1">
                        {registerForm.formState.errors.email.message}
                      </p>
                    )}
                  </div>

                  <div>
                    <label htmlFor="phone" className="block mb-2.5">
                      เบอร์โทรศัพท์ <span className="text-red">*</span>
                    </label>

                    <input
                      {...registerForm.register('phone')}
                      type="tel"
                      id="phone"
                      placeholder="กรอกเบอร์โทรศัพท์"
                      className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                    />
                    {registerForm.formState.errors.phone && (
                      <p className="text-red-500 text-sm mt-1">
                        {registerForm.formState.errors.phone.message}
                      </p>
                    )}
                  </div>
                </div>

                {/* Third Row - Password Fields */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-5 mb-5">
                  <div>
                    <label htmlFor="password" className="block mb-2.5">
                      รหัสผ่าน <span className="text-red">*</span>
                    </label>

                    <input
                      {...registerForm.register('password')}
                      type="password"
                      id="password"
                      placeholder="กรอกรหัสผ่าน"
                      autoComplete="new-password"
                      className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                    />
                    {registerForm.formState.errors.password && (
                      <p className="text-red-500 text-sm mt-1">
                        {registerForm.formState.errors.password.message}
                      </p>
                    )}
                  </div>

                  <div>
                    <label htmlFor="confirmPassword" className="block mb-2.5">
                      ยืนยันรหัสผ่าน <span className="text-red">*</span>
                    </label>

                    <input
                      {...registerForm.register('confirmPassword')}
                      type="password"
                      id="confirmPassword"
                      placeholder="ยืนยันรหัสผ่าน"
                      autoComplete="new-password"
                      className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                    />
                    {registerForm.formState.errors.confirmPassword && (
                      <p className="text-red-500 text-sm mt-1">
                        {registerForm.formState.errors.confirmPassword.message}
                      </p>
                    )}
                  </div>
                </div>

                {/* Fourth Row - Address (Full Width) */}
                <div className="mb-5">
                  <label htmlFor="address" className="block mb-2.5">
                    ที่อยู่ <span className="text-red">*</span>
                  </label>

                  <textarea
                    {...registerForm.register('address')}
                    id="address"
                    placeholder="กรอกที่อยู่"
                    rows={3}
                    className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20 resize-none"
                  />
                  {registerForm.formState.errors.address && (
                    <p className="text-red-500 text-sm mt-1">
                      {registerForm.formState.errors.address.message}
                    </p>
                  )}
                </div>

                <button
                  type="submit"
                  disabled={isLoading}
                  className="w-full flex justify-center font-medium text-white bg-dark py-3 px-6 rounded-lg ease-out duration-200 hover:bg-blue mt-7.5 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {isLoading ? 'กำลังสมัครสมาชิก...' : 'สมัครสมาชิก'}
                </button>

                <p className="text-center mt-6">
                  มีบัญชีผู้ใช้แล้ว?
                  <Link
                    href="/auth/signin"
                    className="text-dark ease-out duration-200 hover:text-blue pl-2"
                  >
                    เข้าสู่ระบบทันที!
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

export default Signup
