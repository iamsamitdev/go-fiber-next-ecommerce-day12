export type Product = {
  id: string
  title: string
  price: number
  discountedPrice: number
  reviews?: number
  description?: string
  stock?: number
  category_id?: string
  imgs?: {
    thumbnails: string[]
    previews: string[]
  }
}