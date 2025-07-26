'use client'

import React, { useState } from 'react'
import { useAuth } from '@/hooks/useAuth'
import { useForm, SubmitHandler } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'

// Zod Schema for Personal Info
const personalInfoSchema = z.object({
  first_name: z.string().min(2, 'ชื่อต้องมีอย่างน้อย 2 ตัวอักษร'),
  last_name: z.string().min(2, 'นามสกุลต้องมีอย่างน้อย 2 ตัวอักษร'),
  email: z.string().email('รูปแบบอีเมลไม่ถูกต้อง'),
  phone: z.string().optional(),
  address: z.string().optional()
})
type PersonalInfoFormData = z.infer<typeof personalInfoSchema>

// Zod Schema for Security
const securitySchema = z.object({
  currentPassword: z.string().min(6, 'รหัสผ่านปัจจุบันต้องมีอย่างน้อย 6 ตัวอักษร'),
  newPassword: z.string().min(6, 'รหัสผ่านใหม่ต้องมีอย่างน้อย 6 ตัวอักษร'),
  confirmNewPassword: z.string()
}).refine(data => data.newPassword === data.confirmNewPassword, {
  message: 'รหัสผ่านใหม่ไม่ตรงกัน',
  path: ['confirmNewPassword']
})
type SecurityFormData = z.infer<typeof securitySchema>

const ProfileComponent = () => {
  const { user, isLoading: authLoading } = useAuth()
  const [activeTab, setActiveTab] = useState('personal')
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [message, setMessage] = useState<{ type: 'success' | 'error', text: string } | null>(null)

  const personalInfoForm = useForm<PersonalInfoFormData>({
    resolver: zodResolver(personalInfoSchema),
    defaultValues: {
      first_name: user?.first_name || '',
      last_name: user?.last_name || '',
      email: user?.email || '',
      phone: user?.phone || '',
      address: user?.address || ''
    }
  })

  const securityForm = useForm<SecurityFormData>({
    resolver: zodResolver(securitySchema)
  })
  
  // Update default values when user data is loaded
  React.useEffect(() => {
    if (user) {
      personalInfoForm.reset({
        first_name: user.first_name,
        last_name: user.last_name,
        email: user.email,
        phone: user.phone,
        address: user.address,
      })
    }
  }, [user, personalInfoForm])

  const onPersonalInfoSubmit: SubmitHandler<PersonalInfoFormData> = async (data) => {
    setIsSubmitting(true)
    setMessage(null)
    console.log('Updating personal info:', data)
    // TODO: Implement API call to update personal info
    await new Promise(resolve => setTimeout(resolve, 1000)); // Simulate API call
    setMessage({ type: 'success', text: 'อัปเดตข้อมูลส่วนตัวสำเร็จ' })
    setIsSubmitting(false)
  }

  const onSecuritySubmit: SubmitHandler<SecurityFormData> = async (data) => {
    setIsSubmitting(true)
    setMessage(null)
    console.log('Updating password')
    // TODO: Implement API call to update password
    await new Promise(resolve => setTimeout(resolve, 1000)); // Simulate API call
    setMessage({ type: 'success', text: 'เปลี่ยนรหัสผ่านสำเร็จ' })
    securityForm.reset()
    setIsSubmitting(false)
  }

  if (authLoading) {
    return (
      <div className="p-6">
        <div className="flex items-center justify-center h-96">
          <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-600"></div>
        </div>
      </div>
    )
  }

  return (
    <div className="p-6">
      <div className="max-w-4xl mx-auto">
        {/* Profile Header */}
        <div className="bg-white rounded-xl shadow-sm border p-6 mb-8">
          <div className="flex flex-col sm:flex-row items-center gap-6">
            <div className="relative">
              <div className="w-24 h-24 bg-blue-100 rounded-full flex items-center justify-center text-blue-600 text-4xl font-bold">
                {user?.first_name?.charAt(0) || 'A'}
              </div>
              <button className="absolute bottom-0 right-0 w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center shadow hover:bg-blue-700 transition-colors">
                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </button>
            </div>
            <div>
              <h1 className="text-2xl font-bold text-gray-900">{user?.first_name} {user?.last_name}</h1>
              <p className="text-gray-600">{user?.role?.name === 'admin' ? 'ผู้ดูแลระบบ' : 'ผู้ใช้งาน'}</p>
            </div>
          </div>
        </div>

        {/* Tabs */}
        <div className="mb-6 border-b border-gray-200">
          <nav className="flex space-x-6" aria-label="Tabs">
            <button
              onClick={() => setActiveTab('personal')}
              className={`px-3 py-2 font-medium text-sm rounded-t-lg transition-colors ${
                activeTab === 'personal'
                  ? 'border-b-2 border-blue-600 text-blue-600'
                  : 'text-gray-500 hover:text-gray-700'
              }`}
            >
              ข้อมูลส่วนตัว
            </button>
            <button
              onClick={() => setActiveTab('security')}
              className={`px-3 py-2 font-medium text-sm rounded-t-lg transition-colors ${
                activeTab === 'security'
                  ? 'border-b-2 border-blue-600 text-blue-600'
                  : 'text-gray-500 hover:text-gray-700'
              }`}
            >
              ความปลอดภัย
            </button>
          </nav>
        </div>
        
        {message && (
          <div className={`mb-5 p-4 rounded-lg text-center ${
            message.type === 'success' 
              ? 'bg-green-50 text-green-700 border border-green-200' 
              : 'bg-red-50 text-red-700 border border-red-200'
          }`}>
            {message.text}
          </div>
        )}

        {/* Tab Content */}
        <div className="bg-white rounded-xl shadow-sm border p-6">
          {activeTab === 'personal' && (
            <form onSubmit={personalInfoForm.handleSubmit(onPersonalInfoSubmit)}>
              <h2 className="text-lg font-semibold text-gray-900 mb-6">แก้ไขข้อมูลส่วนตัว</h2>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label htmlFor="first_name" className="block text-sm font-medium text-gray-700 mb-1">ชื่อจริง</label>
                  <input {...personalInfoForm.register('first_name')} type="text" id="first_name" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                  {personalInfoForm.formState.errors.first_name && <p className="text-red-500 text-sm mt-1">{personalInfoForm.formState.errors.first_name.message}</p>}
                </div>
                <div>
                  <label htmlFor="last_name" className="block text-sm font-medium text-gray-700 mb-1">นามสกุล</label>
                  <input {...personalInfoForm.register('last_name')} type="text" id="last_name" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                  {personalInfoForm.formState.errors.last_name && <p className="text-red-500 text-sm mt-1">{personalInfoForm.formState.errors.last_name.message}</p>}
                </div>
                <div className="md:col-span-2">
                  <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">อีเมล</label>
                  <input {...personalInfoForm.register('email')} type="email" id="email" className="w-full px-4 py-2 border border-gray-300 rounded-lg bg-gray-100" readOnly />
                </div>
                <div>
                  <label htmlFor="phone" className="block text-sm font-medium text-gray-700 mb-1">เบอร์โทรศัพท์</label>
                  <input {...personalInfoForm.register('phone')} type="text" id="phone" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                </div>
                <div className="md:col-span-2">
                  <label htmlFor="address" className="block text-sm font-medium text-gray-700 mb-1">ที่อยู่</label>
                  <textarea {...personalInfoForm.register('address')} id="address" rows={3} className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500"></textarea>
                </div>
              </div>
              <div className="mt-6 text-right">
                <button type="submit" disabled={isSubmitting} className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
                  {isSubmitting ? 'กำลังบันทึก...' : 'บันทึกการเปลี่ยนแปลง'}
                </button>
              </div>
            </form>
          )}

          {activeTab === 'security' && (
            <form onSubmit={securityForm.handleSubmit(onSecuritySubmit)}>
              <h2 className="text-lg font-semibold text-gray-900 mb-6">เปลี่ยนรหัสผ่าน</h2>
              <div className="space-y-6">
                <div>
                  <label htmlFor="currentPassword" className="block text-sm font-medium text-gray-700 mb-1">รหัสผ่านปัจจุบัน</label>
                  <input {...securityForm.register('currentPassword')} type="password" id="currentPassword" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                  {securityForm.formState.errors.currentPassword && <p className="text-red-500 text-sm mt-1">{securityForm.formState.errors.currentPassword.message}</p>}
                </div>
                <div>
                  <label htmlFor="newPassword" className="block text-sm font-medium text-gray-700 mb-1">รหัสผ่านใหม่</label>
                  <input {...securityForm.register('newPassword')} type="password" id="newPassword" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                  {securityForm.formState.errors.newPassword && <p className="text-red-500 text-sm mt-1">{securityForm.formState.errors.newPassword.message}</p>}
                </div>
                <div>
                  <label htmlFor="confirmNewPassword" className="block text-sm font-medium text-gray-700 mb-1">ยืนยันรหัสผ่านใหม่</label>
                  <input {...securityForm.register('confirmNewPassword')} type="password" id="confirmNewPassword" className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500" />
                  {securityForm.formState.errors.confirmNewPassword && <p className="text-red-500 text-sm mt-1">{securityForm.formState.errors.confirmNewPassword.message}</p>}
                </div>
              </div>
              <div className="mt-6 text-right">
                <button type="submit" disabled={isSubmitting} className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
                  {isSubmitting ? 'กำลังเปลี่ยนรหัสผ่าน...' : 'เปลี่ยนรหัสผ่าน'}
                </button>
              </div>
            </form>
          )}
        </div>
      </div>
    </div>
  )
}

export default ProfileComponent
