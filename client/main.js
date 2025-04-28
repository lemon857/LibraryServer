const address = "http://localhost:8080";

let template;
let container;

let banner;

let is_server_available;

let is_cached_books;
let cached_books;

function SetBannerVisibility(show) {
  banner.style.visibility = show ? "visible" : "hidden";
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
    container.appendChild(clone);
  });
}

function ErrorHandler(err) {
  console.log(err.Error());
}

function UnavailableServerHandler(err) {
  is_server_available = false;
  SetBannerVisibility(true);
}

function SyncBooksList() {
  fetch(address + "/books", { method: "GET" }).then(response => response.json()).then(BooksUpdate).catch(ErrorHandler);
}

function CheckHealthServer() {
  is_server_available = true;
  SetBannerVisibility(false);
  fetch(address + "/health", { method: "GET" }).catch(UnavailableServerHandler);
}

document.addEventListener('DOMContentLoaded', function() {
  template = document.getElementById('book-template');
  container = document.getElementById('book-list');
  banner = document.getElementById('banner');

  is_cached_books = false;

  CheckHealthServer();

  SyncBooksList();
});
