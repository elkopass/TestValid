import { defineFeature, loadFeature } from 'jest-cucumber';
import { getDocument } from 'pptr-testing-library';

const feature = loadFeature('./features/add.feature');

defineFeature(feature, test => {
    test('Результат валидации скобок', ({ given, when, then }) => {
        given(/^Страница открыта "(.*)"$/, async (url) => {
            await page.goto(url)
            await page.waitForTimeout(3000);
        });

        when(/^Отправка формы с введеным текстом "(.*)"$/, async (text) => {
            const $document = await getDocument(page);
            await page.type('input[id=input_text]', text)
            await page.waitForTimeout(1000);
            await page.click('#btn_save');
        });

        then(/^Получения результата валидации$/, async (text) => {
            await page.waitForTimeout(3000);
        });
    });


});