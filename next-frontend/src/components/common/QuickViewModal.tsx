"use client"

import React, { useEffect, useState } from "react"
import { useModalContext } from "@/context/QuickViewModalContext"
import Image from "next/image"
import useCartStore from "@/stores/cartStore"
import { formatPrice } from "@/lib/utils"
import { useCartModalContext } from "@/context/CartSidebarModalContext"
import useWishlistStore from "@/stores/wishlistStore"

const QuickViewModal = () => {
  const { isModalOpen, closeModal, product } = useModalContext()
  const { addToCart } = useCartStore()
  const { openCartModal } = useCartModalContext()
  const { addToWishlist } = useWishlistStore()
  const [quantity, setQuantity] = useState(1)
  const [activePreview, setActivePreview] = useState(0)

  const handleAddToCart = () => {
    if (!product) return
    const productToAdd = { ...product, quantity }
    addToCart(productToAdd)
    closeModal()
    openCartModal()
  }

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      const target = event.target as Element
      if (!target.closest(".modal-content")) {
        closeModal()
      }
    }

    if (isModalOpen) {
      document.addEventListener("mousedown", handleClickOutside)
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside)
      setQuantity(1)
      setActivePreview(0)
    }
  }, [isModalOpen, closeModal])

  if (!product) {
    return null
  }

  return (
    <div
      className={`${
        isModalOpen ? "z-99999" : "hidden"
      } fixed top-0 left-0 overflow-y-auto no-scrollbar w-full h-screen sm:py-20 xl:py-25 2xl:py-[230px] bg-dark/70 sm:px-8 px-4 py-5`}
    >
      <div className="flex items-center justify-center ">
        <div className="w-full max-w-[1100px] rounded-xl shadow-3 bg-white p-7.5 relative modal-content">
          <button
            onClick={() => closeModal()}
            aria-label="button for close modal"
            className="absolute top-0 right-0 sm:top-6 sm:right-6 flex items-center justify-center w-10 h-10 rounded-full ease-in duration-150 bg-meta text-body hover:text-dark"
          >
            <svg
              className="fill-current"
              width="26"
              height="26"
              viewBox="0 0 26 26"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M14.3108 13L19.2291 8.08167C19.5866 7.72417 19.5866 7.12833 19.2291 6.77083C19.0543 6.59895 18.8189 6.50262 18.5737 6.50262C18.3285 6.50262 18.0932 6.59895 17.9183 6.77083L13 11.6892L8.08164 6.77083C7.90679 6.59895 7.67142 6.50262 7.42623 6.50262C7.18104 6.50262 6.94566 6.59895 6.77081 6.77083C6.41331 7.12833 6.41331 7.72417 6.77081 8.08167L11.6891 13L6.77081 17.9183C6.41331 18.2758 6.41331 18.8717 6.77081 19.2292C7.12831 19.5867 7.72414 19.5867 8.08164 19.2292L13 14.3108L17.9183 19.2292C18.2758 19.5867 18.8716 19.5867 19.2291 19.2292C19.5866 18.8717 19.5866 18.2758 19.2291 17.9183L14.3108 13Z"
                fill=""
              />
            </svg>
          </button>

          <div className="flex flex-wrap items-center gap-12.5">
            <div className="max-w-[526px] w-full">
              <div className="flex gap-5">
                <div className="flex flex-col gap-5">
                  {product.imgs?.thumbnails?.map((img, key) => (
                    <button
                      onClick={() => setActivePreview(key)}
                      key={key}
                      className={`relative flex items-center justify-center w-20 h-20 overflow-hidden rounded-lg bg-gray-1 ease-out duration-200 hover:border-2 hover:border-blue ${
                        activePreview === key && "border-2 border-blue"
                      }`}
                    >
                      <Image
                        src={img || "/images/products/product-1-sm-1.png"}
                        alt="thumbnail"
                        fill
                        className="object-cover"
                      />
                    </button>
                  ))}
                </div>

                <div className="relative z-1 overflow-hidden flex items-center justify-center w-full sm:min-h-[508px] bg-gray-1 rounded-lg border border-gray-3">
                  <div>
                    <button
                      onClick={() => {}}
                      aria-label="button for zoom"
                      className="gallery__Image w-10 h-10 rounded-[5px] bg-white shadow-1 flex items-center justify-center ease-out duration-200 text-dark hover:text-blue absolute top-4 lg:top-8 right-4 lg:right-8 z-50"
                    >
                      <svg
                        className="fill-current"
                        width="22"
                        height="22"
                        viewBox="0 0 22 22"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          fillRule="evenodd"
                          clipRule="evenodd"
                          d="M9.11493 1.14581L9.16665 1.14581C9.54634 1.14581 9.85415 1.45362 9.85415 1.83331C9.85415 2.21301 9.54634 2.52081 9.16665 2.52081C7.41873 2.52081 6.17695 2.52227 5.23492 2.64893C4.31268 2.77292 3.78133 3.00545 3.39339 3.39339C3.00545 3.78133 2.77292 4.31268 2.64893 5.23492C2.52227 6.17695 2.52081 7.41873 2.52081 9.16665C2.52081 9.54634 2.21301 9.85415 1.83331 9.85415C1.45362 9.85415 1.14581 9.54634 1.14581 9.16665L1.14581 9.11493C1.1458 7.43032 1.14579 6.09599 1.28619 5.05171C1.43068 3.97699 1.73512 3.10712 2.42112 2.42112C3.10712 1.73512 3.97699 1.43068 5.05171 1.28619C6.09599 1.14579 7.43032 1.1458 9.11493 1.14581ZM16.765 2.64893C15.823 2.52227 14.5812 2.52081 12.8333 2.52081C12.4536 2.52081 12.1458 2.21301 12.1458 1.83331C12.1458 1.45362 12.4536 1.14581 12.8333 1.14581L12.885 1.14581C14.5696 1.1458 15.904 1.14579 16.9483 1.28619C18.023 1.43068 18.8928 1.73512 19.5788 2.42112C20.2648 3.10712 20.5693 3.97699 20.7138 5.05171C20.8542 6.09599 20.8542 7.43032 20.8541 9.11494V9.16665C20.8541 9.54634 20.5463 9.85415 20.1666 9.85415C19.787 9.85415 19.4791 9.54634 19.4791 9.16665C19.4791 7.41873 19.4777 6.17695 19.351 5.23492C19.227 4.31268 18.9945 3.78133 18.6066 3.39339C18.2186 3.00545 17.6873 2.77292 16.765 2.64893ZM1.83331 12.1458C2.21301 12.1458 2.52081 12.4536 2.52081 12.8333C2.52081 14.5812 2.52227 15.823 2.64893 16.765C2.77292 17.6873 3.00545 18.2186 3.39339 18.6066C3.78133 18.9945 4.31268 19.227 5.23492 19.351C6.17695 19.4777 7.41873 19.4791 9.16665 19.4791C9.54634 19.4791 9.85415 19.787 9.85415 20.1666C9.85415 20.5463 9.54634 20.8541 9.16665 20.8541H9.11494C7.43032 20.8542 6.09599 20.8542 5.05171 20.7138C3.97699 20.5693 3.10712 20.2648 2.42112 19.5788C1.73512 18.8928 1.43068 18.023 1.28619 16.9483C1.14579 15.904 1.1458 14.5696 1.14581 12.885L1.14581 12.8333C1.14581 12.4536 1.45362 12.1458 1.83331 12.1458ZM20.1666 12.1458C20.5463 12.1458 20.8541 12.4536 20.8541 12.8333V12.885C20.8542 14.5696 20.8542 15.904 20.7138 16.9483C20.5693 18.023 20.2648 18.8928 19.5788 19.5788C18.8928 20.2648 18.023 20.5693 16.9483 20.7138C15.904 20.8542 14.5696 20.8542 12.885 20.8541H12.8333C12.4536 20.8541 12.1458 20.5463 12.1458 20.1666C12.1458 19.787 12.4536 19.4791 12.8333 19.4791C14.5812 19.4791 15.823 19.4777 16.765 19.351C17.6873 19.227 18.2186 18.9945 18.6066 18.6066C18.9945 18.2186 19.227 17.6873 19.351 16.765C19.4777 15.823 19.4791 14.5812 19.4791 12.8333C19.4791 12.4536 19.787 12.1458 20.1666 12.1458Z"
                          fill=""
                        />
                      </svg>
                    </button>

                    {product?.imgs?.previews?.[activePreview] ? (
                      <Image
                        src={product.imgs.previews[activePreview]}
                        alt="products-details"
                        fill
                        className="object-cover"
                      />
                    ) : (
                      <div className="w-[400px] h-[400px] bg-gray-200 flex items-center justify-center">
                        <span className="text-gray-500">ไม่มีรูปภาพ</span>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            </div>

            <div className="max-w-[445px] w-full">
              <span
                className={`inline-block text-custom-xs font-medium text-white py-1 px-3 mb-6.5 ${
                  product.stock && product.stock > 0 ? "bg-green" : "bg-red"
                }`}
              >
                {product.stock && product.stock > 0
                  ? "In Stock"
                  : "Out of Stock"}
              </span>

              <h3 className="font-semibold text-xl xl:text-heading-5 text-dark mb-4">
                {product.title}
              </h3>

              <div className="flex flex-wrap items-center gap-5 mb-6">
                <div className="flex items-center gap-1.5">
                  {/* <!-- stars --> */}
                  <div className="flex items-center gap-1">
                    <svg
                      className="fill-[#FFA645]"
                      width="18"
                      height="18"
                      viewBox="0 0 18 18"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clipPath="url(#clip0_375_9172)">
                        <path
                          d="M16.7906 6.72187L11.7 5.93438L9.39377 1.09688C9.22502 0.759375 8.77502 0.759375 8.60627 1.09688L6.30002 5.9625L1.23752 6.72187C0.871891 6.77812 0.731266 7.25625 1.01252 7.50938L4.69689 11.3063L3.82502 16.6219C3.76877 16.9875 4.13439 17.2969 4.47189 17.0719L9.05627 14.5687L13.6125 17.0719C13.9219 17.2406 14.3156 16.9594 14.2313 16.6219L13.3594 11.3063L17.0438 7.50938C17.2688 7.25625 17.1563 6.77812 16.7906 6.72187Z"
                          fill=""
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_375_9172">
                          <rect width="18" height="18" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>

                    <svg
                      className="fill-[#FFA645]"
                      width="18"
                      height="18"
                      viewBox="0 0 18 18"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clipPath="url(#clip0_375_9172)">
                        <path
                          d="M16.7906 6.72187L11.7 5.93438L9.39377 1.09688C9.22502 0.759375 8.77502 0.759375 8.60627 1.09688L6.30002 5.9625L1.23752 6.72187C0.871891 6.77812 0.731266 7.25625 1.01252 7.50938L4.69689 11.3063L3.82502 16.6219C3.76877 16.9875 4.13439 17.2969 4.47189 17.0719L9.05627 14.5687L13.6125 17.0719C13.9219 17.2406 14.3156 16.9594 14.2313 16.6219L13.3594 11.3063L17.0438 7.50938C17.2688 7.25625 17.1563 6.77812 16.7906 6.72187Z"
                          fill=""
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_375_9172">
                          <rect width="18" height="18" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>

                    <svg
                      className="fill-[#FFA645]"
                      width="18"
                      height="18"
                      viewBox="0 0 18 18"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clipPath="url(#clip0_375_9172)">
                        <path
                          d="M16.7906 6.72187L11.7 5.93438L9.39377 1.09688C9.22502 0.759375 8.77502 0.759375 8.60627 1.09688L6.30002 5.9625L1.23752 6.72187C0.871891 6.77812 0.731266 7.25625 1.01252 7.50938L4.69689 11.3063L3.82502 16.6219C3.76877 16.9875 4.13439 17.2969 4.47189 17.0719L9.05627 14.5687L13.6125 17.0719C13.9219 17.2406 14.3156 16.9594 14.2313 16.6219L13.3594 11.3063L17.0438 7.50938C17.2688 7.25625 17.1563 6.77812 16.7906 6.72187Z"
                          fill=""
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_375_9172">
                          <rect width="18" height="18" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>

                    <svg
                      className="fill-gray-4"
                      width="18"
                      height="18"
                      viewBox="0 0 18 18"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clipPath="url(#clip0_375_9172)">
                        <path
                          d="M16.7906 6.72187L11.7 5.93438L9.39377 1.09688C9.22502 0.759375 8.77502 0.759375 8.60627 1.09688L6.30002 5.9625L1.23752 6.72187C0.871891 6.77812 0.731266 7.25625 1.01252 7.50938L4.69689 11.3063L3.82502 16.6219C3.76877 16.9875 4.13439 17.2969 4.47189 17.0719L9.05627 14.5687L13.6125 17.0719C13.9219 17.2406 14.3156 16.9594 14.2313 16.6219L13.3594 11.3063L17.0438 7.50938C17.2688 7.25625 17.1563 6.77812 16.7906 6.72187Z"
                          fill=""
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_375_9172">
                          <rect width="18" height="18" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>

                    <svg
                      className="fill-gray-4"
                      width="18"
                      height="18"
                      viewBox="0 0 18 18"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clipPath="url(#clip0_375_9172)">
                        <path
                          d="M16.7906 6.72187L11.7 5.93438L9.39377 1.09688C9.22502 0.759375 8.77502 0.759375 8.60627 1.09688L6.30002 5.9625L1.23752 6.72187C0.871891 6.77812 0.731266 7.25625 1.01252 7.50938L4.69689 11.3063L3.82502 16.6219C3.76877 16.9875 4.13439 17.2969 4.47189 17.0719L9.05627 14.5687L13.6125 17.0719C13.9219 17.2406 14.3156 16.9594 14.2313 16.6219L13.3594 11.3063L17.0438 7.50938C17.2688 7.25625 17.1563 6.77812 16.7906 6.72187Z"
                          fill=""
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_375_9172">
                          <rect width="18" height="18" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>
                  </div>
                  <p className="font-medium text-dark-4">
                    4.7 Rating ({product.reviews || 0} reviews)
                  </p>
                </div>
                <span className="w-px h-5 bg-gray-4 block"></span>
                <div className="flex items-center gap-2">
                  <svg
                    width="20"
                    height="20"
                    viewBox="0 0 20 20"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <g clipPath="url(#clip0_375_9185)">
                      <path
                        d="M9.99984 1.66666C5.4165 1.66666 1.6665 5.41666 1.6665 10C1.6665 14.5833 5.4165 18.3333 9.99984 18.3333C14.5832 18.3333 18.3332 14.5833 18.3332 10C18.3332 5.41666 14.5832 1.66666 9.99984 1.66666ZM9.99984 16.6667C6.3165 16.6667 3.33317 13.6833 3.33317 10C3.33317 6.31666 6.3165 3.33332 9.99984 3.33332C13.6832 3.33332 16.6665 6.31666 16.6665 10C16.6665 13.6833 13.6832 16.6667 9.99984 16.6667Z"
                        fill="#16B676"
                      />
                      <path
                        d="M13.2416 7.20832L9.31659 11.1333C9.17492 11.275 8.94992 11.275 8.80826 11.1333L6.75826 9.08332C6.61659 8.94166 6.61659 8.71666 6.75826 8.57499C6.90009 8.43332 7.12492 8.43332 7.26659 8.57499L9.06659 10.375L12.7333 6.70832C12.875 6.56666 13.0999 6.56666 13.2416 6.70832C13.3833 6.84999 13.3833 7.07499 13.2416 7.20832Z"
                        fill="#16B676"
                      />
                    </g>
                    <defs>
                      <clipPath id="clip0_375_9185">
                        <rect width="20" height="20" fill="white" />
                      </clipPath>
                    </defs>
                  </svg>
                  <p className="font-medium text-dark">
                    {product.stock && product.stock > 0
                      ? "In Stock"
                      : "Out of Stock"}
                  </p>
                </div>
              </div>

              <p className="font-medium text-dark-5 mb-6">{product.description}</p>

              <div className="flex items-center gap-12.5 mb-11">
                <div className="flex flex-col">
                  <span className="font-medium text-dark-4 mb-1">ราคา</span>
                  <span className="font-medium text-2xl text-dark">
                    {formatPrice(product.price)}
                  </span>
                </div>

                <div className="flex flex-col">
                  <span className="font-medium text-dark-4 mb-2.5">จำนวน</span>
                  <div className="flex items-center rounded-md border border-gray-3 w-max">
                    <button
                      onClick={() => setQuantity(quantity > 1 ? quantity - 1 : 1)}
                      aria-label="button for remove product"
                      className="flex items-center justify-center w-11.5 h-11.5 ease-out duration-200 hover:text-blue"
                    >
                      <svg
                        className="fill-current"
                        width="16"
                        height="2"
                        viewBox="0 0 16 2"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          fillRule="evenodd"
                          clipRule="evenodd"
                          d="M-8.548e-08 0.977778C-3.82707e-08 0.437766 0.437766 3.82707e-08 0.977778 8.548e-08L15.0222 1.31328e-06C15.5622 1.36049e-06 16 0.437767 16 0.977779C16 1.51779 15.5622 1.95556 15.0222 1.95556L0.977778 1.95556C0.437766 1.95556 -1.32689e-07 1.51779 -8.548e-08 0.977778Z"
                          fill=""
                        />
                      </svg>
                    </button>
                    <span className="flex items-center justify-center w-16 h-11.5 border-x border-gray-4">
                      {quantity}
                    </span>
                    <button
                      onClick={() => setQuantity(quantity + 1)}
                      aria-label="button for add product"
                      className="flex items-center justify-center w-11.5 h-11.5 ease-out duration-200 hover:text-blue"
                    >
                      <svg
                        className="fill-current"
                        width="16"
                        height="16"
                        viewBox="0 0 16 16"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          fillRule="evenodd"
                          clipRule="evenodd"
                          d="M8.08889 0C8.6289 2.36047e-08 9.06667 0.437766 9.06667 0.977778L9.06667 15.0222C9.06667 15.5622 8.6289 16 8.08889 16C7.54888 16 7.11111 15.5622 7.11111 15.0222L7.11111 0.977778C7.11111 0.437766 7.54888 -2.36047e-08 8.08889 0Z"
                          fill=""
                        />
                        <path
                          fillRule="evenodd"
                          clipRule="evenodd"
                          d="M0 7.91111C4.72093e-08 7.3711 0.437766 6.93333 0.977778 6.93333L15.0222 6.93333C15.5622 6.93333 16 7.3711 16 7.91111C16 8.45112 15.5622 8.88889 15.0222 8.88889L0.977778 8.88889C0.437766 8.88889 -4.72093e-08 8.45112 0 7.91111Z"
                          fill=""
                        />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>

              <div className="flex items-center gap-4">
                <button
                  onClick={handleAddToCart}
                  aria-label="button for add to cart"
                  className="inline-flex font-medium text-white bg-blue py-3 px-8 rounded-md ease-out duration-200 hover:bg-blue-dark"
                >
                  เพิ่มลงตะกร้า
                </button>
                <button
                  onClick={() => product && addToWishlist(product)}
                  aria-label="button for add to wishlist"
                  className="inline-flex font-medium text-dark bg-gray-1 py-3 px-8 rounded-md ease-out duration-200 hover:bg-dark hover:text-white"
                >
                  เพิ่มลงรายการที่ชอบ
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default QuickViewModal
