<template>
    <div>
      <Navigasi />
      <main class="main-teams">
        <h1 class="title">All Teams</h1>
        <div class="teams-grid">
          <div v-for="team in teams" :key="team.id" class="team-card">
            <h2>{{ team.nama }}</h2>
            <p>Kelas: {{ team.kelas }}</p>
            <h3>Pembalap:</h3>
            <ul>
              <li v-for="rider in team.pembalap" :key="rider.id">
                {{ rider.nama }} (#{{ rider.nomer }})
              </li>
            </ul>
          </div>
        </div>
      </main>
      <Footer />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const teams = ref([])

const fetchTeams = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/v1/tim')
    teams.value = response.data
  } catch (error) {
    console.error('Error fetching teams:', error)
  }
}

onMounted(fetchTeams)
</script>

<style scoped>
.main-teams {
  margin-inline: 100px;
  margin-top: 40px;
}

.title {
  font-family: 'MotoGP Display', sans-serif;
  font-weight: bold;
  font-size: 40px;
  margin-bottom: 60px;
}

.teams-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 30px;
}

.team-card {
  background-color: #f5f5f5;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

@media (max-width: 768px) {
  .main-teams {
    margin-inline: 20px;
  }

  .teams-grid {
    grid-template-columns: 1fr;
  }
}
</style>