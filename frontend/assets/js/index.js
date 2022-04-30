const preloader = document.querySelector`#preloader`
const booksTable = document.querySelector`#books-table`
const noBooks = document.querySelector`#no-books`

const loadBooks = async () => {
    preloader.removeAttribute`style`
    noBooks.style.display = 'none'
    booksTable.style.display = 'none'
    booksTable.querySelector`tbody`.innerHTML = ''

    const response = await fetch('http://localhost:7777/book')

    const json = await response.json()

    preloader.style.display = 'none'

    if (!json || !json.length) {
        noBooks.removeAttribute`style`
    } else {
        json.forEach(book => {
            booksTable.querySelector`tbody`.innerHTML += `
                <tr>
                    <td>${book.Id}</td>
                    <td>
                        <input type="text" name="Name" id="Name" value="${book.Name}" data-id="${book.Id}" class="book-name" />
                    </td>
                    <td>
                        <input type="number" name="Cost" id="Cost" step="0.1" value="${book.Cost}" data-id="${book.Id}" class="book-cost" />
                    </td>
                    <td>
                        <button class="update-book" data-id="${book.Id}">Update</button>
                        <button class="destroy-book" data-id="${book.Id}">Destroy</button>
                    </td>
                </tr>
            `
        })

        booksTable.removeAttribute`style`
        
        document.querySelectorAll`.update-book`.forEach(btn => {
            btn.addEventListener('click', () => {
                const id = btn.dataset.id

                const name = document.querySelector(`.book-name[data-id="${id}"]`).value
                const cost = document.querySelector(`.book-cost[data-id="${id}"]`).value

                const formData = new FormData()
                formData.append('Id', id)
                formData.append('Name', name)
                formData.append('Cost', cost)

                updateBook(formData)
            })
        })

        document.querySelectorAll`.destroy-book`.forEach(btn => {
            btn.addEventListener('click', () => {
                const id = btn.dataset.id

                const formData = new FormData()
                formData.append('Id', id)

                destroyBook(formData)
            })
        })
    }
}

loadBooks()

const storeBookForm = document.querySelector`#store-book`

storeBookForm.addEventListener('submit', async e => {
    e.preventDefault()

    const formData = new FormData(e.target)

    storeBook(formData)
})

const storeBook = async formData => {
    await fetch('http://localhost:7777/book/store', {
        method: 'post',
        body: formData,
    })

    loadBooks()
}

const updateBook = async formData => {
    await fetch('http://localhost:7777/book/update', {
        method: 'post',
        body: formData,
    })

    loadBooks()
}

const destroyBook = async formData => {
    await fetch('http://localhost:7777/book/destroy', {
        method: 'post',
        body: formData,
    })

    loadBooks()
}
