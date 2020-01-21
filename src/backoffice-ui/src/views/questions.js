import m from "mithril";
import level from "../components/bulma/level";

export default function Questions(initialVnode) {
    const { talkId } = initialVnode.attrs;

    let questions = [];
    let favMap = {};
    let showOnlyFav = false;

    function loadQuestions() {
        return m.request({
            method: 'GET',
            url: `/api/talks/${talkId}/questions`,
        }).then((result) => {
            questions = result;
        });
    }

    function getVisibleQuestions() {
        if (!showOnlyFav) return questions;
        return questions.filter(q => favMap[q.id.toString()])
    }

    function readFavMap() {
        favMap = JSON.parse(localStorage.getItem(`${talkId}-fav-map`) || '{}')
    }

    function writeFavMap() {
        localStorage.setItem(`${talkId}-fav-map`, JSON.stringify(favMap))
    }

    function toggleFav(question) {
        favMap[question.id.toString()] = !favMap[question.id.toString()]
        writeFavMap()
    }

    function init() {
        loadQuestions();
        readFavMap();
    }

    return {
        view: () => m('div', [
            level(
                m('h1', { class: 'title is-3' }, 'Questions'),
                m('label.checkbox', { style: 'user-select: none;' }, [
                    m('input', { type: 'checkbox', style: 'margin-right:8px;', checked: showOnlyFav, onchange: (e) => showOnlyFav = e.target.checked }),
                    'Show only favorite'
                ])
            ),
            m('div.table-container', [
                m('table', { class: 'table is-fullwidth is-striped' }, [
                    m('thead', [
                        m('tr', [
                            m('th', 'Question'),
                            m('th', 'Favourite'),
                        ]),
                    ]),
                    m('tbody', getVisibleQuestions().map(question => m('tr', { key: question.id }, [
                        m('td', question.question),
                        m('td', [
                            m('button.button.is-small', {
                                class: favMap[question.id.toString()] ? 'is-warning' : '',
                                onclick: () => toggleFav(question)
                            }, 'Fav')
                        ])
                    ]))),
                ])
            ]),
        ]),
        oninit: () => init(),
    };
}
