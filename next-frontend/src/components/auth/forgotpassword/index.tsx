import Link from "next/link"
import React from "react"

const ForgotPassword = () => {
  return (
    <>
      <section className="flex items-center justify-center">
        <div className="w-full max-w-md">
          <div className="bg-white rounded-lg shadow-lg p-8">
            <div className="text-center mb-8">
              <h2 className="font-semibold text-2xl text-dark mb-2">
                ลืมรหัสผ่าน
              </h2>
              <p className="text-gray-600">กรอกอีเมลเพื่อรีเซ็ตรหัสผ่าน</p>
            </div>

            <div>
              <form>
                <div className="mb-5">
                  <label htmlFor="email" className="block mb-2.5">
                    อีเมล
                  </label>

                  <input
                    type="email"
                    name="email"
                    id="email"
                    placeholder="กรอกอีเมลของคุณ"
                    className="rounded-lg border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full py-3 px-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                  />
                </div>

                <button
                  type="submit"
                  className="w-full flex justify-center font-medium text-white bg-dark py-3 px-6 rounded-lg ease-out duration-200 hover:bg-blue mt-7.5"
                >
                  ส่งลิงก์รีเซ็ตรหัสผ่าน
                </button>

                <div className="text-center mt-6 space-y-2">
                  <p>
                    จำรหัสผ่านได้แล้ว?
                    <Link
                      href="/auth/signin"
                      className="text-dark ease-out duration-200 hover:text-blue pl-2"
                    >
                      เข้าสู่ระบบ
                    </Link>
                  </p>
                  <p>
                    ยังไม่มีบัญชี?
                    <Link
                      href="/auth/signup"
                      className="text-dark ease-out duration-200 hover:text-blue pl-2"
                    >
                      สมัครสมาชิกทันที!
                    </Link>
                  </p>
                </div>
              </form>
            </div>
          </div>
        </div>
      </section>
    </>
  )
}

export default ForgotPassword
