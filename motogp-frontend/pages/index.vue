<template>
    <div>
      <h1>Daftar Tim MotoGP</h1>
      <ul>
        <li v-for="tim in tims" :key="tim.id">
          {{ tim.nama }} ({{ tim.pembalap }})
        </li>
      </ul>
      <p v-if="loading">Loading...</p>
      <p v-if="error">{{ error }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const tims = ref([])
  const loading = ref(true)
  const error = ref(null)
  
  const fetchTim = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/v1/tim')
      tims.value = response.data
    } catch (err) {
      console.error('Error fetching tim:', err)
      error.value = 'Gagal memuat data tim.'
    } finally {
      loading.value = false
    }
  }
  
  onMounted(() => {
    fetchTim()
  })
  </script>
  