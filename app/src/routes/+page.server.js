export const load = async () => {
    const getPlaces = async () => {
        const res = await fetch(`http://127.0.0.1:8080/get`)
        const data = await res.json()
        return data
    }

    return {
        places: await getPlaces(),
    }
}