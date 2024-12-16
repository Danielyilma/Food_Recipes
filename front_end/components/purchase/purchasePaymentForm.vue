<template>
    <div class="p-6">
      <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-6">Payment Details</h2>
      
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- <div class="space-y-4">
          <FormInput
            id="cardName"
            label="Name on Card"
            v-model="formData.cardName"
            placeholder="Enter the name on your card"
            required
          />
          
          <FormInput
            id="cardNumber"
            label="Card Number"
            v-model="formData.cardNumber"
            placeholder="1234 5678 9012 3456"
            required
            @input="formatCardNumber"
          />
          
          <div class="grid grid-cols-2 gap-4">
            <FormInput
              id="expiryDate"
              label="Expiry Date"
              v-model="formData.expiryDate"
              placeholder="MM/YY"
              required
              @input="formatExpiryDate"
            />
            
            <FormInput
              id="cvv"
              label="CVV"
              v-model="formData.cvv"
              type="password"
              placeholder="123"
              required
              maxlength="4"
            />
          </div>
        </div>        -->
        <!-- Order Summary -->
        <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
          <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Order Summary</h3>
          <div class="flex justify-between text-sm text-gray-600 dark:text-gray-300 mb-2">
            <span>Subtotal</span>
            <span>${{ amount.toFixed(2) }}</span>
          </div>
          <!-- <div class="flex justify-between text-sm text-gray-600 dark:text-gray-300 mb-4">
            <span>Tax</span>
            <span>${{ (amount * 0.1).toFixed(2) }}</span>
          </div> -->
          <div class="flex justify-between text-lg font-medium text-gray-900 dark:text-white">
            <span>Total</span>
            <span>${{ (amount).toFixed(2) }}</span>
          </div>
        </div>
        
        <!-- Actions -->
        <div class="flex justify-between m-2">
          <button
            type="button"
            @click="$emit('back')"
            class="px-6 py-2 border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700"
          >
            Back
          </button>
          
          <button
            type="submit"
            class="px-6 py-2 bg-primary-600 text-white rounded-lg disabled:opacity-50 hover:bg-primary-700"
          >
            Pay Now
          </button>
        </div>
      </form>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue'
  
  const props = defineProps<{
    amount: number
  }>()
  
  const emit = defineEmits<{
    (e: 'back'): void
    (e: 'submit', paymentDetails: any): void
  }>()
  
  const formData = ref({
    cardName: '',
    cardNumber: '',
    expiryDate: '',
    cvv: ''
  })
  
  const formatCardNumber = (event: Event) => {
    const input = event.target as HTMLInputElement
    let value = input.value.replace(/\D/g, '')
    value = value.substring(0, 16)
    value = value.replace(/(\d{4})(?=\d)/g, '$1 ')
    formData.value.cardNumber = value
  }
  
  const formatExpiryDate = (event: Event) => {
    const input = event.target as HTMLInputElement
    let value = input.value.replace(/\D/g, '')
    value = value.substring(0, 4)
    if (value.length > 2) {
      value = value.substring(0, 2) + '/' + value.substring(2)
    }
    formData.value.expiryDate = value
  }
  
  const isValid = computed(() => {
    return (
      formData.value.cardName.length > 0 &&
      formData.value.cardNumber.replace(/\s/g, '').length === 16 &&
      formData.value.expiryDate.length === 5 &&
      formData.value.cvv.length >= 3
    )
  })
  
  const handleSubmit = () => {
    emit('submit', {
      ...formData.value,
      amount: props.amount
    })
  }
  </script>