document.addEventListener('DOMContentLoaded', function() {
    const imageContainer = document.getElementById('imageContainer');
    const images = document.querySelectorAll('#imageContainer img');
    let idx = 0;

    document.getElementById('nextBtn').addEventListener('click', () => {
        if (idx < images.length - 1) {
            idx++;
            updateImagePosition();
        }
    });

    document.getElementById('prevBtn').addEventListener('click', () => {
        if (idx > 0) {
            idx--;
            updateImagePosition();
        }
    });

    function updateImagePosition() {
        imageContainer.style.transform = `translateX(${-idx * 100}%)`;
    }
});
