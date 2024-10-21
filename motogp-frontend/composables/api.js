// composables/api.js
import axios from 'axios'

// Base URL API backend Golang
const API_BASE_URL = 'http://localhost:8080/api/v1'

// Fetch daftar tim
export const fetchTim = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/tim`)
    return response.data
  } catch (error) {
    console.error('Error fetching tim:', error)
    return []
  }
}

// Fetch pembalap berdasarkan tim ID
export const fetchPembalapByTimId = async (timId) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/tim/${timId}`)
    return response.data.Pembalap
  } catch (error) {
    console.error('Error fetching pembalap:', error)
    return []
  }
}
