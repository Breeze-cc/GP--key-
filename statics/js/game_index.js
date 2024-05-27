function performSearch() {
    const query = document.getElementById('searchInput').value;
    const resultsDiv = document.getElementById('results');

    // Clear previous results
    resultsDiv.innerHTML = '';

    // Mock search results
    const mockResults = [
        { title: '老奶奶环', description: 'Description.' },
        { title: '失败的man', description: 'Description' },
        { title: '大表哥2', description: 'Description' },
    ];

    mockResults.forEach(result => {
        const resultItem = document.createElement('div');
        resultItem.className = 'result-item';

        const resultTitle = document.createElement('h3');
        resultTitle.textContent = result.title;

        const resultDescription = document.createElement('p');
        resultDescription.textContent = result.description;

        resultItem.appendChild(resultTitle);
        resultItem.appendChild(resultDescription);

        resultsDiv.appendChild(resultItem);
    });
}
