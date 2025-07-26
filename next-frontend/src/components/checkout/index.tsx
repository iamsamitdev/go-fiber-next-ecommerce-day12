"use client"
import React, { useState, useEffect } from "react"
import Breadcrumb from "@/components/common/Breadcrumb"
import Login from "@/components/checkout/Login"
import Shipping from "@/components/checkout/Shipping"
import ShippingMethod from "@/components/checkout/ShippingMethod"
import PaymentMethod from "@/components/checkout/PaymentMethod"
import Coupon from "@/components/checkout/Coupon"
import Billing from "@/components/checkout/Billing"
import useCartStore from "@/stores/cartStore"
import { formatPrice } from "@/lib/utils"
import { useAuth } from "@/hooks/useAuth"

const Checkout = () => {
  const { cartItems, cartTotal } = useCartStore()
  const { isAuthenticated } = useAuth()
  const [isClient, setIsClient] = useState(false)
  const shippingCost = 15 // Placeholder for shipping cost

  useEffect(() => {
    setIsClient(true)
  }, [])

  const finalTotal = cartTotal() + shippingCost

  return (
    <>
      <Breadcrumb title={"เช็คเอาท์"} pages={["เช็คเอาท์"]} />
      <section className="overflow-hidden py-20 bg-gray-2">
        <div className="max-w-[1170px] w-full mx-auto px-4 sm:px-8 xl:px-0">
          <form>
            <div className="flex flex-col lg:flex-row gap-7.5 xl:gap-11">
              {/* <!-- checkout left --> */}
              <div className="lg:max-w-[670px] w-full">
                {!isAuthenticated && (
                  <>
                    {/* <!-- login box --> */}
                    <Login />

                    {/* <!-- billing details --> */}
                    <Billing />
                  </>
                )}

                {/* <!-- address box two --> */}
                <Shipping />

                {/* <!-- others note box --> */}
                <div className="bg-white shadow-1 rounded-[10px] p-4 sm:p-8.5 mt-7.5">
                  <div>
                    <label htmlFor="notes" className="block mb-2.5">
                      บันทึกข้อมูลอื่นๆ (ไม่บันทึกได้)
                    </label>

                    <textarea
                      name="notes"
                      id="notes"
                      rows={5}
                      placeholder="บันทึกข้อมูลอื่นๆ เช่น ข้อมูลการจัดส่ง"
                      className="rounded-md border border-gray-3 bg-gray-1 placeholder:text-dark-5 w-full p-5 outline-none duration-200 focus:border-transparent focus:shadow-input focus:ring-2 focus:ring-blue/20"
                    ></textarea>
                  </div>
                </div>
              </div>

              {/* // <!-- checkout right --> */}
              <div className="max-w-[455px] w-full">
                {/* <!-- order list box --> */}
                <div className="bg-white shadow-1 rounded-[10px]">
                  <div className="border-b border-gray-3 py-5 px-4 sm:px-8.5">
                    <h3 className="font-medium text-xl text-dark">
                      รายการสินค้า
                    </h3>
                  </div>

                  <div className="pt-2.5 pb-8.5 px-4 sm:px-8.5">
                    {/* <!-- title --> */}
                    <div className="flex items-center justify-between py-5 border-b border-gray-3">
                      <div>
                        <h4 className="font-medium text-dark">สินค้า</h4>
                      </div>
                      <div>
                        <h4 className="font-medium text-dark text-right">
                          ราคาสินค้า
                        </h4>
                      </div>
                    </div>

                    {isClient &&
                      cartItems.map((item) => (
                        <div
                          key={item.id}
                          className="flex items-center justify-between py-5 border-b border-gray-3"
                        >
                          <div>
                            <p className="text-dark">
                              {item.title} x {item.quantity}
                            </p>
                          </div>
                          <div>
                            <p className="text-dark text-right">
                              {formatPrice(item.price * item.quantity)}
                            </p>
                          </div>
                        </div>
                      ))}

                    {/* <!-- shipping cost --> */}
                    <div className="flex items-center justify-between py-5 border-b border-gray-3">
                      <div>
                        <p className="text-dark">ค่าจัดส่ง</p>
                      </div>
                      <div>
                        <p className="text-dark text-right">
                          {isClient ? formatPrice(shippingCost) : formatPrice(0)}
                        </p>
                      </div>
                    </div>

                    {/* <!-- total --> */}
                    <div className="flex items-center justify-between pt-5">
                      <div>
                        <p className="font-medium text-lg text-dark">ราคารวม</p>
                      </div>
                      <div>
                        <p className="font-medium text-lg text-dark text-right">
                          {isClient ? formatPrice(finalTotal) : formatPrice(0)}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>

                {/* <!-- coupon box --> */}
                <Coupon />

                {/* <!-- shipping box --> */}
                <ShippingMethod />

                {/* <!-- payment box --> */}
                <PaymentMethod />

                {/* <!-- checkout button --> */}
                <button
                  type="submit"
                  className="w-full flex justify-center font-medium text-white bg-blue py-3 px-6 rounded-md ease-out duration-200 hover:bg-blue-dark mt-7.5"
                >
                  ดำเนินการชำระเงิน
                </button>
              </div>
            </div>
          </form>
        </div>
      </section>
    </>
  )
}

export default Checkout
