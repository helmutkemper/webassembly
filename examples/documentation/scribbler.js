// utilities
var get = function (selector, scope) {
  scope = scope ? scope : document;
  return scope.querySelector(selector);
};

var getAll = function (selector, scope) {
  scope = scope ? scope : document;
  return scope.querySelectorAll(selector);
};

// setup typewriter effect in the terminal demo
if (document.getElementsByClassName('demo').length > 0) {
  var i = 0;
  var txt1 = `cd ./examples/server; make build
            `;
  var txt2 = `2022/05/04 12:08:50 addr: fe80::9d3:9afa:56f6:69a5
2022/05/04 12:08:50 addr: 169.254.105.165
2022/05/04 12:08:50 addr: fe80::b8af:7b2d:1006:4c20
2022/05/04 12:08:50 addr: 169.254.76.32
2022/05/04 12:08:50 addr: fe80::e929:7524:6b9b:78ab
2022/05/04 12:08:50 addr: 169.254.120.171
2022/05/04 12:08:50 addr: fe80::a44a:1de9:f2de:83f1
2022/05/04 12:08:50 addr: 192.168.25.29
2022/05/04 12:08:50 addr: ::1
2022/05/04 12:08:50 addr: 127.0.0.1
2022/05/04 12:08:50 addr: fe80::4168:f357:d716:a3a7
2022/05/04 12:08:50 addr: 172.27.208.1
2022/05/04 12:08:50 Listening on :3000...
`;
  var speed = 80;
  var txt = txt1
  var tipping = 0



  function typeItOut () {
    if (i < txt.length) {
      document.getElementsByClassName('demo')[0].innerHTML += txt.charAt(i);
      i++;
      setTimeout(typeItOut, speed);
    } else if (tipping === 0) {
      tipping = 1;
      typeItOut();
    } else if (tipping === 1) {
      tipping = false;
      txt = txt2;
      i = 0;
      speed = 0;
      setTimeout(typeItOut, speed);
    }
  }

  setTimeout(typeItOut, 1800);
}

// toggle tabs on codeblock
window.addEventListener("load", function() {
  // get all tab_containers in the document
  var tabContainers = getAll(".tab__container");

  // bind click event to each tab container
  for (var i = 0; i < tabContainers.length; i++) {
    get('.tab__menu', tabContainers[i]).addEventListener("click", tabClick);
  }

  // each click event is scoped to the tab_container
  function tabClick (event) {
    var scope = event.currentTarget.parentNode;
    var clickedTab = event.target;
    var tabs = getAll('.tab', scope);
    var panes = getAll('.tab__pane', scope);
    var activePane = get(`.${clickedTab.getAttribute('data-tab')}`, scope);

    // remove all active tab classes
    for (var i = 0; i < tabs.length; i++) {
      tabs[i].classList.remove('active');
    }

    // remove all active pane classes
    for (var i = 0; i < panes.length; i++) {
      panes[i].classList.remove('active');
    }

    // apply active classes on desired tab and pane
    clickedTab.classList.add('active');
    activePane.classList.add('active');
  }
});

//in page scrolling for documentaiton page
var btns = getAll('.js-btn');
var sections = getAll('.js-section');

function setActiveLink(event) {
  // remove all active tab classes
  for (var i = 0; i < btns.length; i++) {
    btns[i].classList.remove('selected');
  }

  event.target.classList.add('selected');
}

function smoothScrollTo(i, event) {
  var element = sections[i];
  setActiveLink(event);

  window.scrollTo({
    'behavior': 'smooth',
    'top': element.offsetTop - 20,
    'left': 0
  });
}

if (btns.length && sections.length > 0) {
  for (var i = 0; i<btns.length; i++) {
    btns[i].addEventListener('click', smoothScrollTo.bind(this,i));
  }
}

// fix menu to page-top once user starts scrolling
window.addEventListener('scroll', function () {
  var docNav = get('.doc__nav > ul');

  if( docNav) {
    if (window.pageYOffset > 63) {
      docNav.classList.add('fixed');
    } else {
      docNav.classList.remove('fixed');
    }
  }
});

// responsive navigation
var topNav = get('.menu');
var icon = get('.toggle');

window.addEventListener('load', function(){
  function showNav() {
    if (topNav.className === 'menu') {
      topNav.className += ' responsive';
      icon.className += ' open';
    } else {
      topNav.className = 'menu';
      icon.classList.remove('open');
    }
  }
  icon.addEventListener('click', showNav);
});

