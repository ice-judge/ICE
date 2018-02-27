# 문제 관련 api

## /problems GET

### 설명

문제들의 리스트를 불러옵니다.

### 쿼리

| 필드 | 설명 | 값 |
| --- | --- | --- |
| orderby  | 정렬기준 | string |
| dir | 정렬방향(오름차순 내림차순) | bool |
| difficulty | 난이도 | string |
| tag | 태그 | string |
| p | 페이지 | int |

### 출력

| 필드 | 설명 | 값 |
| --- | --- | --- |
| problems | 문제의 간략한 정보들 | [문제헤더](../schema/problems.md)[] |
| prev_pages | 이전 페이지들을의 api 주소 | 페이지링크[] |
| next_pages | 다음 페이지들을의 api 주소 | 페이지링크[] |


## /difficulties GET

### 설명

모든 난이도의 리스트를 불러옵니다.

### 출력

| 필드 | 설명 | 값 |
| --- | --- | --- |
| difficulties | 태그 | 난이도[] |

## /tags

### 설명

모든 태그들의 리스트를 불러옵니다

### 쿼리

| 필드 | 설명 | 값 |
| --- | --- | --- |
| orderby  | 정렬기준 | string |
| dir | 정렬방향(오름차순 내림차순) | bool |
| p | 페이지 | int |

### 출력

| 필드 | 설명 | 값 |
| --- | --- | --- |
| tags | 태그 | 태그[] |
