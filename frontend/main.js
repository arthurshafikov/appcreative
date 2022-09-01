(function(){
    const getCurrentWeatherUrl = 'http://localhost:8123/v1/getCurrentWeather'

    document.querySelector('form').addEventListener('submit', (e) => {
        e.preventDefault()

        const searchBtn = e.target.querySelector('button[type="submit"]')
        searchBtn.disabled = true
        searchBtn.classList.add('loading')

        const errorBox = document.querySelector('.errors')
        errorBox.classList.add('hidden')

        fetch(getCurrentWeatherUrl + '?city=' + e.target.querySelector('input').value)
            .then(response => {
                return response.json()
            })
            .then(json => {
                if (json.errors !== undefined) {
                    const errors = json.errors
                    for (const name in errors) {
                        errorBox.innerHTML = errors[name].join('\n')
                    }
                    errorBox.classList.remove('hidden')
                }

                document.querySelector('#city-name').innerHTML = json?.city
                document.querySelector('#temperature').innerHTML = json?.temperature
                document.querySelector('#weather-conditions').innerHTML = json?.weatherCondition?.type
                document.querySelector('#wind-speed').innerHTML = json?.wind?.speed
                document.querySelector('#wind-direction').innerHTML = json?.wind?.direction
                document.querySelector('#pressure').innerHTML = json?.weatherCondition?.pressure
                document.querySelector('#humidity').innerHTML = json?.weatherCondition?.humidity


                searchBtn.disabled = false
                searchBtn.classList.remove('loading')
            })
    })
})()
