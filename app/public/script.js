document.getElementById('fetchData').addEventListener('click', fetchData)

function fetchData() {
  fetch('/api/v1/members')
    .then(response => response.json())
    .then(data => {
      const apiDataElement = document.getElementById('apiData')
      apiDataElement.innerHTML = ''
      const memberList = document.createElement('ul')

      data.forEach(member => {
        const listItem = document.createElement('li')
        listItem.innerHTML = `<strong>ID:</strong> ${member.id}<br><strong>Name:</strong> ${member.name}`
        listItem.classList.add('member-item')
        memberList.appendChild(listItem)
      })

      apiDataElement.appendChild(memberList)
    })
    .catch(error => {
      console.error('Error fetching data:', error)
    })
}
