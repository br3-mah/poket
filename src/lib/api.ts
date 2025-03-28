import axios from 'axios';

const API = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
});


// Add response interceptor for error handling
API.interceptors.response.use(
  response => response,
  error => {
    if (error.response) {
      // Server responded with non-2xx status
      const { status, data } = error.response;
      
      if (status === 500) {
        console.error('Server Error:', data);
        return Promise.reject({
          message: 'Internal server error',
          details: data.details || 'unknown_error'
        });
      }
      
      return Promise.reject(data);
    }
    return Promise.reject(error);
  }
);

export const registerUser = async (userData: {
  name: string;
  email: string;
  password: string;
}) => {
  try {
    const response = await API.post('/register', userData);
    return response.data;
  } catch (error) {
    console.error('Registration Error:', error);
    throw error; // Re-throw for component handling
  }
};

export const loginUser = async (credentials: {
  email: string;
  password: string;
}) => {
  return API.post('/login', credentials);
};