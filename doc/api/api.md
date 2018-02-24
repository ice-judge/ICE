# Document

## /problems/{difficulty} GET

### 설명

문제들의 리스트를 불러옵니다.

### URL

{difficulty} 난이도 만약 all이면 모든 난이도

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향(오름차순 내림차순)

{p} 페이지

### 에러코드


### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]

## /problems/{difficulty}/search GET

### 설명

카테고리에서 문제를 검색합니다.

## URL

{category} 카테고리

### 쿼리

{q} 키워드

{section} 검색할 부분

{orderby} 정렬기준

{dir} 정렬방향

{p} 페이지

### 출력

problems:문제[]

prev_pages:페이지링크[]

next_pages:페이지링크[]

## /category GET

### 설명

모든 카테고리의 리스트를 불러옵니다.

### 쿼리

{orderby} 정렬기준

{dir} 정렬방향

### 에러코드

### 출력

catergories:카테고리[]

## / GET
