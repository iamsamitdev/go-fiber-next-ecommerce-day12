'use client'

import { useState, useEffect } from 'react'
import axios from 'axios'

export interface Category {
  id: string
  name: string
  description: string
  image: string
}

interface ApiResponse {
  success: boolean
  message: string
  data: Category[]
  pagination: {
    page: number
    limit: number
    total_pages: number
    total_items: number
  }
}

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const useCategory = () => {
  const [categories, setCategories] = useState<Category[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchCategories = async () => {
      setIsLoading(true)
      setError(null)
      try {
        const response = await api.get<ApiResponse>('/api/v1/categories') // Fetch all categories
        if (response.data.success) {
          setCategories(response.data.data)
        } else {
          setError(response.data.message || 'Failed to fetch categories')
        }
      } catch (err: any) {
        console.error('Fetch categories error:', err)
        setError(err.response?.data?.message || 'An error occurred while fetching categories.')
      } finally {
        setIsLoading(false)
      }
    }

    fetchCategories()
  }, [])

  return { categories, isLoading, error }
} 