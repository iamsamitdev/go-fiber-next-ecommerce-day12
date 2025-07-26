import { Menu } from "@/types/Menu"

export const menuData: Menu[] = [
  {
    id: 1,
    title: "ขายดี",
    newTab: false,
    path: "/",
  },
  {
    id: 2,
    title: "ร้านค้า",
    newTab: false,
    path: "/shop-with-sidebar",
  },
  {
    id: 3,
    title: "ติดต่อ",
    newTab: false,
    path: "/contact",
  },
  {
    id: 4,
    title: "เพจ",
    newTab: false,
    path: "/",
    submenu: [
      {
        id: 41,
        title: "เช็คเอาท์",
        newTab: false,
        path: "/checkout",
      },
      {
        id: 42,
        title: "ตะกร้า",
        newTab: false,
        path: "/cart",
      },
      {
        id: 43,
        title: "รายการสิ่งที่อยากได้",
        newTab: false,
        path: "/wishlist",
      },
    ],
  },
  {
    id: 5,
    title: "บล็อก",
    newTab: false,
    path: "/",
    submenu: [
      {
        id: 51,
        title: "บล็อกกริด",
        newTab: false,
        path: "/blogs/blog-grid-with-sidebar",
      },
      {
        id: 52,
        title: "บล็อกรายละเอียด",
        newTab: false,
        path: "/blogs/blog-details",
      },
    ],
  },
]
