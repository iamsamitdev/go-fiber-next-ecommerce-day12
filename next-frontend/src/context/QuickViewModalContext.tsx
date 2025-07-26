"use client"

import { Product } from "@/types/product"
import React, { createContext, useContext, useState } from "react"

interface ModalContextType {
  isModalOpen: boolean
  product: Product | null
  openModal: (product: Product) => void
  closeModal: () => void
}

const ModalContext = createContext<ModalContextType | undefined>(undefined)

export const useModalContext = () => {
  const context = useContext(ModalContext)
  if (!context) {
    throw new Error("useModalContext must be used within a ModalProvider")
  }
  return context
}

export const ModalProvider = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [product, setProduct] = useState<Product | null>(null)

  const openModal = (selectedProduct: Product) => {
    setProduct(selectedProduct)
    setIsModalOpen(true)
  }

  const closeModal = () => {
    setIsModalOpen(false)
    setProduct(null)
  }

  return (
    <ModalContext.Provider
      value={{ isModalOpen, openModal, closeModal, product }}
    >
      {children}
    </ModalContext.Provider>
  )
} 