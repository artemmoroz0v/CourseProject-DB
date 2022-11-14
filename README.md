# База данных фитнес-клуба

## Интерфейс

### Клиент фитнес-клуба

- [x] Добавление клиента
- [ ] Удаление клиента
- [x] Список клиентов
- [x] Список неподписанных клиентов
- [ ] Изменение роста/веса клиента

### Тренер

- [x] Добавление тренера
- [ ] Удаление тренера
- [x] Список тренеров

### Группа фитнес-клуба

- [x] Добавление группы
- [ ] Удаление группы
- [x] Список всех групп
- [ ] Список одной группы
- [ ] Добавление клиента в группу
- [ ] Удаление клиента из группы

### Расписание

- [ ] Вывод расписания для каждой группы
- [ ] Вывод расписания для каждой программы
- [ ] Добавление записей в расписание
- [ ] Удаление записей из расписания


## Схема базы данных

![image](https://github.com/papey08/DB/blob/main/DBscheme.png)


## Таблицы

### Client

Содержит основную информацию о всех клиентах фитнес-клуба

### Trainer

Содержит основную информацию о всех тренерах фитнес-клуба

### Group

Содержит информацию о всех группах фитнес-клуба: тренер, программу группы и 
в какое время занимаются

### Group_client

Содержит информацию, в какой группе какие клиенты занимаются. **Тип связи 
многие ко многим** (подразумевается, что в группе может быть много клиентов, и 
один клиент может быть в нескольких группах)

### Program

Содержит список программ фитнес-клуба (для каждой программы есть свой 
определённый зал, **тип связи один к одному**, поэтому таблицу для залов мы не 
делаем)

### Times

Содержит список окон, в которые проводятся занятия.

### Timetable

Содержит расписание групп фитнес-клуба, **тип связи многие ко многим**. Одна 
группа может занимать много окон, в одно окно могут заниматься многие группы, 
так как зал у двух произвольных групп может быть разным за счёт того, что 
программы могут быть разные. **Важно:** за счёт того, что поле time_id 
уникальное, невозможен вариант, когда у двух групп, занимающихся в одном 
зале, стоят одинаковые время и день занятия.
