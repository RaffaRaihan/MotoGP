<template>
    <div>
      <Navigasi />
      <main class="main-riders">
        <h1 class="title">All Riders</h1>
        <div class="btn-kelas">
          <button @click="filterRiders('MotoGP')" class="btn-motogp">MotoGP</button>
          <button @click="filterRiders('Moto2')" class="btn-motogp">Moto2</button>
          <button @click="filterRiders('Moto3')" class="btn-motogp">Moto3</button>
        </div>
        <div class="card">
          <div v-for="rider in filteredRiders" :key="rider.id" class="card-riders">
            <div class="hastag">
              <img :src="rider.gambar" :alt="rider.nama">
              <h1>#{{ rider.nomer }}</h1>
            </div>
            <div class="profile">
              <h1>{{ rider.nama }}</h1>
              <p>{{ rider.tim.nama }}</p>
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const riders = ref([])
const filteredRiders = ref([])
const currentClass = ref('MotoGP')

const fetchRiders = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/v1/pembalap')
    riders.value = response.data
    filterRiders(currentClass.value)
  } catch (error) {
    console.error('Error fetching riders:', error)
  }
}

const filterRiders = (kelas) => {
  currentClass.value = kelas
  filteredRiders.value = riders.value.filter(rider => rider.tim.kelas === kelas)
}

onMounted(fetchRiders)
</script>

<style scoped>
/* main riders start */
.main-riders{
    margin-inline: 100px;
    margin-top: 40px;
}

.main-riders .title{
    font-family: 'MotoGP Display';
    font-weight: bold;
    font-size: 40px;
    margin-bottom: 60px;
}
.btn-kelas{
    display: flex;
    gap: 25px;
}

.btn-motogp{
    background: white;
    border: solid;
    color: #000;
    border-radius: 15px;
    text-decoration: none;
    font-family: 'MotoGP Display';
    font-weight: normal;
    font-size: 32px;
    padding: 13px 21px;
}
.btn-motogp :hover{
    background: #CD1818;
    color: #000;
    font-family: 'MotoGP Display';
    font-weight: normal;
    font-size: 32px;
    text-decoration: underline;
}

.card{
    margin-top: 70px;
    display: flex;
    justify-content: space-between;
    gap: 74px;
}

.card-riders{
    width: 254px;
    height: 392px;
    background: rgb(0,0,0);
    background: linear-gradient(360deg, rgba(0,0,0,1) 0%, rgba(195,0,0,1) 100%);
    border-radius: 10px;
}
.hastag{
    display: flex;
}
/* main riders end */

/* Tambahkan CSS responsif jika diperlukan */
@media (max-width: 768px) {
  .card {
    flex-direction: column;
  }
  .card-riders {
    width: 100%;
    margin-bottom: 20px;
  }
}
</style>