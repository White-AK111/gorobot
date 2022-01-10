# gorobot
Тест robotgo

### Кейс
1. Открыть браузер
2. Отправить текст в окно браузера
3. Отправить enter
4. Переместить мышь
5. Прокрутить scroll
6. Закрыть браузер

### Проблемы
Возможно проблема в том, что Ubuntu запущена на ВМ
1. Ошибка при запуске
Ошибка:
load_input_helper [1883]: XkbGetKeyboard failed to locate a valid keyboard!
2. Не получается снять изображение с экрана
Ошибка:
X Error of failed request:  BadMatch (invalid parameter attributes)
  Major opcode of failed request:  73 (X_GetImage)
  Serial number of failed request:  7
  Current serial number in output stream:  7
exit status 1
3. Функция MoveSmooth не передвигает курсор мыши
