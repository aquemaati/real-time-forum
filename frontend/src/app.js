const posts = [
    { 
        id: 1, 
        title: "How to Perform SQL Injection", 
        user: "cyber_warrior", 
        date: "2024-06-01", 
        description: "A comprehensive guide to performing SQL Injection attacks, including techniques and prevention methods.", 
        likes: 120, 
        dislikes: 5, 
        category: "Web Application Testing" 
    },
    { 
        id: 2, 
        title: "Cracking WPA2 WiFi Passwords", 
        user: "wifi_cracker", 
        date: "2024-05-28", 
        description: "Step-by-step tutorial on cracking WPA2 WiFi passwords using popular tools.", 
        likes: 150, 
        dislikes: 12, 
        category: "Wireless Network Testing" 
    },
    { 
        id: 3, 
        title: "Introduction to Metasploit", 
        user: "exploit_master", 
        date: "2024-05-30", 
        description: "Learn how to use Metasploit for various types of penetration testing.", 
        likes: 200, 
        dislikes: 3, 
        category: "Tools and Software" 
    },
    { 
        id: 4, 
        title: "Phishing Attack Simulation", 
        user: "social_engineer", 
        date: "2024-06-02", 
        description: "Simulating a phishing attack to test security awareness within an organization.", 
        likes: 95, 
        dislikes: 8, 
        category: "Social Engineering" 
    },
    { 
        id: 5, 
        title: "Exploit Development for Beginners", 
        user: "code_infiltrator", 
        date: "2024-05-25", 
        description: "An introductory guide to writing your own exploits.", 
        likes: 130, 
        dislikes: 7, 
        category: "Vulnerability Research and Exploits" 
    },
    { 
        id: 6, 
        title: "Burp Suite for Web Pen Testing", 
        user: "web_defender", 
        date: "2024-05-27", 
        description: "How to use Burp Suite for effective web application penetration testing.", 
        likes: 110, 
        dislikes: 2, 
        category: "Tools and Software" 
    },
    { 
        id: 7, 
        title: "Defending Against XSS Attacks", 
        user: "security_ninja", 
        date: "2024-06-03", 
        description: "Methods and best practices for defending against Cross-Site Scripting (XSS) attacks.", 
        likes: 140, 
        dislikes: 6, 
        category: "Web Application Testing" 
    },
    { 
        id: 8, 
        title: "Using Wireshark for Network Analysis", 
        user: "packet_sniffer", 
        date: "2024-05-29", 
        description: "A beginner's guide to using Wireshark for analyzing network traffic.", 
        likes: 165, 
        dislikes: 4, 
        category: "Tools and Software" 
    },
    { 
        id: 9, 
        title: "Ethical Hacking: Legal Aspects", 
        user: "legal_hacker", 
        date: "2024-06-04", 
        description: "Understanding the legal considerations and ethics in penetration testing.", 
        likes: 90, 
        dislikes: 1, 
        category: "Ethics and Legal Aspects" 
    },
    { 
        id: 10, 
        title: "Top CTF Challenges for Practice", 
        user: "ctf_champion", 
        date: "2024-06-01", 
        description: "A list of top Capture The Flag (CTF) challenges for practicing your hacking skills.", 
        likes: 175, 
        dislikes: 3, 
        category: "Capture the Flag (CTF) and Challenges" 
    }
];

const routes = [
    { path: '/', view: () => '<h1>Home</h1><p>Welcome to the homepage!</p>' },
    { path: '/about', view: () => '<h1>About</h1><p>About us page.</p>' },
    { path: '/contact', view: () => '<h1>Contact</h1><p>Contact us page.</p>' },
    { path: '/posts', view: () => getPostsView() },
    { path: '/post/:id', view: (params) => getPostView(params) }
];

function getPostsView() {
    return posts.map(post => `
        <div class="post">
            <h2>${post.title}</h2>
            <p><strong>User:</strong> ${post.user}</p>
            <p><strong>Date:</strong> ${post.date}</p>
            <p>${post.description}</p>
            <p><strong>Likes:</strong> ${post.likes} | <strong>Dislikes:</strong> ${post.dislikes}</p>
            <p><strong>Category:</strong> ${post.category}</p>
            <a href="/post/${post.id}" data-link>Read more</a>
        </div>
    `).join('');
}

function getPostView(params) {
    const post = posts.find(post => post.id == params.id);
    if (!post) {
        return '<h1>404 Not Found</h1><p>Post not found.</p>';
    }
    return `
        <div class="post">
            <h1>${post.title}</h1>
            <p><strong>User:</strong> ${post.user}</p>
            <p><strong>Date:</strong> ${post.date}</p>
            <p>${post.description}</p>
            <p><strong>Likes:</strong> ${post.likes} | <strong>Dislikes:</strong> ${post.dislikes}</p>
            <p><strong>Category:</strong> ${post.category}</p>
            <a href="/posts" data-link>Back to posts</a>
        </div>
    `;
}

function navigateTo(url) {
    history.pushState(null, null, url);
    router();
}

function pathToRegex(path) {
    return new RegExp(`^${path.replace(/:\w+/g, '(.+)')}$`);
}

function getParams(match) {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);
    return Object.fromEntries(keys.map((key, i) => [key, values[i]]));
}

function findMatch() {
    const potentialMatches = routes.map(route => ({
        route,
        result: location.pathname.match(pathToRegex(route.path))
    }));

    return potentialMatches.find(potentialMatch => potentialMatch.result !== null) || {
        route: { path: '/404', view: () => '<h1>404 Not Found</h1><p>Page not found.</p>' },
        result: [location.pathname]
    };
}

function router() {
    const app = document.getElementById('app');
    const match = findMatch();
    const params = getParams(match);
    app.innerHTML = match.route.view(params);
}

window.addEventListener('popstate', router);

document.addEventListener('DOMContentLoaded', () => {
    document.body.addEventListener('click', (e) => {
        if (e.target.matches('[data-link]')) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });
    router();
});
