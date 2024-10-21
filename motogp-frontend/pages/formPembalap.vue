<template>
    <div>
      <h1>Tambah Pembalap Baru</h1>
      <form @submit.prevent="submitPembalap">
        <div>
          <label for="nama">Nama Pembalap:</label>
          <input type="text" id="nama" v-model="nama" required />
        </div>
        <div>
          <label for="negara">Negara:</label>
          <input type="text" id="negara" v-model="negara" required />
        </div>
        <div>
          <label for="timId">ID Tim:</label>
          <input type="text" id="timId" v-model="timId" required />
        </div>
        <button type="submit">Tambah Pembalap</button>
      </form>
  
      <p v-if="message">{{ message }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import axios from 'axios'
  
  const nama = ref('')
  const negara = ref('')
  const timId = ref(null)
  const message = ref('')
  
  // Fungsi untuk submit data pembalap ke backend
  const submitPembalap = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/v1/pembalap', {
        Nama: nama.value,
        Negara: negara.value,
        TimID: timId.value
      })
      message.value = 'Pembalap berhasil ditambahkan!'
      nama.value = ''
      negara.value = ''
      timId.value = null // Reset form setelah berhasil
    } catch (error) {
      console.error('Error menambahkan pembalap:', error)
      message.value = 'Gagal menambahkan pembalap!'
    }
  }
  </script>
  