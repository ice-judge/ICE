# Document

## /problems GET

### 설명

모든 문제들의 리스트를 불러옵니다.

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향(오름차순 내림차순)

{p} 페이지

### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]

## /problems/difficulty/{difficulty} GET

### 설명

특정 난이도의 문제들의 리스트를 불러옵니다.

### URL

{difficulty} 난이도

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향(오름차순 내림차순)

{p} 페이지

### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]

## /problems/search GET

### 설명

문제를 검색합니다.

### 쿼리

{q} 키워드 ( {기준}:{키워드}, ... ) ( ex)tag:hoy!,name:hoi! )

{orderby} 정렬기준

{dir} 정렬방향

{p} 페이지

### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]

## /problems/tag/{tag} GET

### 설명

특정 태그의 문제의 리스트를 불러옵니다.

### URL

{tag} 태그 만약 untag면 태그가 없는 문제들

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향(오름차순 내림차순)

{p} 페이지

### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]


## /difficulties GET

### 설명

모든 난이도의 리스트를 불러옵니다.

### 출력

difficulties:난이도[]

## /tags

### 설명

모든 태그들의 리스트를 불러옵니다

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향

{p} 페이지 만약 없다면 첫번째 페이지

### 출력

tags:태그[]

prev_pages:페이지링크[]

next_pages:페이지링크[]
