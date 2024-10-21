<template>
    <div>
      <h1>Tambah Tim Baru</h1>
      <form @submit.prevent="submitTim">
        <div>
          <label for="nama">Nama Tim:</label>
          <input type="text" id="nama" v-model="nama" required />
        </div>
        <button type="submit">Tambah Tim</button>
      </form>
  
      <p v-if="message">{{ message }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import axios from 'axios'
  
  const nama = ref('')
  const message = ref('')
  
  // Fungsi untuk submit data tim ke backend
  const submitTim = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/v1/tim', {
        Nama: nama.value
      })
      message.value = 'Tim berhasil ditambahkan!'
      nama.value = '' // Reset form setelah berhasil
    } catch (error) {
      console.error('Error menambahkan tim:', error)
      message.value = 'Gagal menambahkan tim!'
    }
  }
  </script>
  