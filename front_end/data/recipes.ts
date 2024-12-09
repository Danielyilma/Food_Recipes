import type { Recipe } from "~/types/recipe";

export const recipes: Recipe[] = [
  {
    id: 1,
    title: "Classic Pancakes",
    description: "Fluffy and delicious pancakes perfect for breakfast",
    image:
      "https://images.unsplash.com/photo-1528207776546-365bb710ee93?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    category: "Breakfast",
    prepTime: "20 mins",
    author: {
      name: "John Doe",
      avatar:
        "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
    isPaid: false,
  },
  {
    id: 2,
    title: "Grilled Salmon",
    description: "Healthy and tasty grilled salmon with vegetables",
    image:
      "https://images.unsplash.com/photo-1467003909585-2f8a72700288?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    category: "Main Course",
    prepTime: "35 mins",
    author: {
      name: "Jane Smith",
      avatar:
        "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
    isPaid: true,
  },
  {
    id: 3,
    title: "Chocolate Cake",
    description: "Rich and moist chocolate cake for special occasions",
    image:
      "https://images.unsplash.com/photo-1578985545062-69928b1d9587?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80",
    category: "Desserts",
    prepTime: "75 mins",
    author: {
      name: "Mike Johnson",
      avatar:
        "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
    isPaid: true,
  },
];
