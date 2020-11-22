import { mount } from "@vue/test-utils";
import Article from "@/views/Article.vue";

describe("Article", () => {

    const foundWrapper = mount(Article, {
        props: {
            name: "rest_api"
        }
    });

    const notFoundWrapper = mount(Article, {
        props: {
            name: "api_rest"
        }
    });

    test('article name exist', () => {

        const header = foundWrapper.find("#article");

        setTimeout(() => {
            expect(header.exists()).toBe(true);
        }, 1000);
    });

    test('Not Found does not exist', () => {

        const msg = foundWrapper.find('#not-found');

        setTimeout(() => {
            expect(msg.exists()).toBe(false);
        }, 1000)
    })

    test('article description after fetch', () => {

        const desc = foundWrapper.find("#article-desc");

        setTimeout(() => {
            expect(desc.text()).toBe("REpresentational State Transfer Application Programming Interface");
        }, 1000);
    });

    test('Not found exist', () => {

        const msg = notFoundWrapper.find('#not-found');

        setTimeout(() => {
            expect(msg.exists()).toBe(true);
        }, 1000);
    });

    test('Not found message', () => {

        const msg = notFoundWrapper.find('#not-found');

        setTimeout(() => {
            expect(msg.text()).toBe("No article with this exact name found. Use Edit button in the header to add it");
        }, 1000);
    });
});
