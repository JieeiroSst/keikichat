function newDescription() {
    document.write("title - " + document.title + "<br>");
    document.title = "New title here!";
    document.write("title - " + document.title + "<br>");
    let title_el = document.querySelector("title");
    if (document.title !== title_el) {
        document.title = title_el;
    }
}

$('meta[name="title_el"]').attr("content", newDescription);
