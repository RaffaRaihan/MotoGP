<template>
    <div>
      <h1>Daftar Pembalap MotoGP</h1>
      <ul>
        <li v-for="pembalap in pembalaps" :key="pembalap.id">
          {{ pembalap.nama }} ({{ pembalap.nomer }}), Tim: {{ pembalap.tim.nama }}
        </li>
      </ul>
      <p v-if="loading">Loading...</p>
      <p v-if="error">{{ error }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const pembalaps = ref([])
  const loading = ref(true)
  const error = ref(null)
  
  const fetchPembalap = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/v1/pembalap')
      pembalaps.value = response.data
    } catch (err) {
      console.error('Error fetching pembalap:', err)
      error.value = 'Gagal memuat data pembalap.'
    } finally {
      loading.value = false
    }
  }
  
  onMounted(() => {
    fetchPembalap()
  })
  </script>
  