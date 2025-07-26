'use client'

import { useState, useEffect, useCallback } from 'react'
import axios from 'axios'
import { Product } from '@/types/product'

interface ApiProduct {
  id: string
  name: string
  description: string
  price: number
  stock: number
  images: { image_url: string }[]
  category_id: string
}

interface ApiResponse {
  success: boolean
  message: string
  data: ApiProduct[]
  pagination: {
    page: number
    limit: number
    total_pages: number
    total_items: number
  }
}

// Axios instance
const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const useProduct = () => {
  const [products, setProducts] = useState<Product[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [pagination, setPagination] = useState({
    total: 0,
    page: 1,
    limit: 10,
    last_page: 1
  })

  const fetchProducts = useCallback(
    async (
      page = 1,
      limit = 10,
      sortBy = 'created_at',
      sortOrder = 'desc',
      categoryId: string | null = null,
    ) => {
      setIsLoading(true)
      setError(null)
      try {
        const params = new URLSearchParams({
          page: String(page),
          limit: String(limit),
          sort_by: sortBy,
          sort_order: sortOrder,
        });

        const url = categoryId
          ? `/api/v1/products/category/${categoryId}?${params.toString()}`
          : `/api/v1/products?${params.toString()}`;

        const response = await api.get<ApiResponse>(url);

        if (response.data.success && response.data.data) {
          // Transform API data to frontend format
          const transformedProducts = response.data.data.map((apiProduct) => ({
            id: apiProduct.id,
            title: apiProduct.name,
            price: apiProduct.price,
            discountedPrice: apiProduct.price, // Assuming no discount for now
            description: apiProduct.description,
            stock: apiProduct.stock,
            category_id: apiProduct.category_id,
            reviews: 0, // Placeholder
            imgs: {
              thumbnails: apiProduct.images.map((img) => img.image_url),
              previews: apiProduct.images.map((img) => img.image_url)
            }
          }))

          setProducts(transformedProducts)

          setPagination({
            total: response.data.pagination.total_items,
            page: response.data.pagination.page,
            limit: response.data.pagination.limit,
            last_page: response.data.pagination.total_pages
          })
        } else {
          setError(response.data.message || 'Failed to fetch products')
          setProducts([])
        }
      } catch (err: any) {
        console.error('Fetch products error:', err)
        setError(err.response?.data?.message || 'An error occurred while fetching products.')
        setProducts([])
      } finally {
        setIsLoading(false)
      }
    },
    []
  )

  useEffect(() => {
    fetchProducts()
  }, [fetchProducts])

  return {
    products,
    pagination,
    isLoading,
    error,
    fetchProducts
  }
}
