const address = "http://localhost:8080";

let template;
let container;

let book_plane;

let banner;

let is_server_available;

let is_cached_books;
let cached_books;

function SetBannerVisibility(show) {
  banner.style.visibility = show ? "visible" : "hidden";
}

function BookInfoUpdate(book) {
  book_plane.querySelector('#book-title').textContent = book['title'];
  book_plane.querySelector('#book-description').textContent = book['description'];
  fetch(address + `/authors/${Number(book['authorId'])}`, { method: "GET" }).then(response => response.json()).then(BookInfoUpdateAuthor).catch(ErrorHandler);
}

function BookInfoUpdateAuthor(author) {
  book_plane.querySelector('#book-author').textContent = `${author['lastName']} ${author['firstName']}`;
}

function GetBookInfo(book_id) {
  fetch(address + `/books/${Number(book_id)}`, { method: "GET" }).then(response => response.json()).then(BookInfoUpdate).catch(ErrorHandler);
}

function BooksUpdate(books, force_update = false) {
  if (is_cached_books && !force_update) {
    return;
  }

  is_cached_books = true;
  cached_books = books;

  books.forEach(book => {
    const clone = template.content.cloneNode(true);
    clone.querySelector('#book-title').textContent = book['title'];

    clone.querySelector('#book-button').onclick = function() { GetBookInfo(book['id']); }

    container.appendChild(clone);
  });
}

function ErrorHandler(err) {
  if (err == "TypeError: NetworkError when attempting to fetch resource.") {
    is_server_available = false;
    SetBannerVisibility(true);
  } else {
    console.log(err.Error());
  }
}

function SyncBooksList() {
  fetch(address + "/books", { method: "GET" }).then(response => response.json()).then(BooksUpdate).catch(ErrorHandler);
}

function CheckHealthServer() {
  is_server_available = true;
  SetBannerVisibility(false);
  fetch(address + "/health", { method: "GET" }).catch(ErrorHandler);
}

document.addEventListener('DOMContentLoaded', function() {
  template = document.getElementById('book-template');
  container = document.getElementById('book-list');
  banner = document.getElementById('banner');
  book_plane = document.getElementById('book-plane');

  is_cached_books = false;

  CheckHealthServer();

  SyncBooksList();
});
