# lab-home-office-escala

## Regras
* Não pode ter mais que a metade de funcuionários em home office
* Quando um funcionário estiver de home na sexta, no pode
estar na segunda. Não pode sex-seg ou seg-sex
* Não pode repetir os mesmos dias de home office
* Não pode ter home na sequencia
* Ter que deixar aletório os dias
* Todos os dias, deve ter ao menos um funcionarios por departamento 
* Cada departamento deve ter ao menos 3 funcionarios 
## Como executar a regra 
1. Sortear ordem dos funcionarios por departamento
2. Ordenar descrescente por quantidade de funcionarios por departamento
3. Distribuir sequencialmente pelos dias os funcionarios
4. Aplicar regra 'seg-sex':
    - trocar segunda da próxima semana
5. Aplicar regra: 'sex-seg': 
    - trocar segunda pela terça 
6. Aplicar regra: não sequencial
    - trocar pelo elemento da esquerda ou direita
7. Aplicar regra: não repete dia da semana 
    - trocar pelo elemento da esquerda ou direita
