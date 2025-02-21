const baseURL = "http://localhost:8080"

export const api = {
  get: async (endpoint: string) => {
    const response = await fetch(`${baseURL}${endpoint}`)

    const data = await response.json()
    return data;
  },
  post: async (endpoint: string, body: string) => {
    const response = await fetch(`${baseURL}${endpoint}`, {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body
    })

    const data = await response.json()
    return data;
  },
  put: async (endpoint: string, body: string) => {
    const response = await fetch(`${baseURL}${endpoint}`, {
      method: "PUT",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body
    })

    const data = await response.json()
    return data;
  },
  delete: async (endpoint: string) => {
    const response = await fetch(`${baseURL}${endpoint}`, {
      method: "DELETE",
    })

    const data = await response.json()
    return data;
  }
}