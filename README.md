# Gobot - Ativa Grupo SBF

Bem-vindo ao **Gobot**, o assistente automatizado de atendimento via WhatsApp do Ativa Grupo SBF! Este projeto foi desenvolvido para facilitar o atendimento de colaboradores, ex-colaboradores, parceiros e clientes, oferecendo um menu interativo e respostas automáticas para as principais demandas.

---

## 🚀 Visão Geral

O Gobot é um bot escrito em Go, projetado para operar no WhatsApp, fornecendo um fluxo de atendimento baseado em estágios (stages). Ele permite navegação por menus, execução de handlers específicos para cada etapa e controle de permissões de acesso.

---

## 📋 Funcionalidades

- **Menu Principal Interativo:** Usuários podem navegar facilmente entre opções como Adesão, Aplicativo/Senha, Capital, Empréstimos, Parcerias, Consultoria, Ex-colaborador, Negociação de Dívidas, Informe de Rendimentos, Dúvidas e Encerrar Atendimento.
- **Controle de Stages:** Cada etapa do atendimento é tratada como um "stage", com lógica e permissões próprias.
- **Persistência de Dados:** Utiliza banco de dados para armazenar o progresso e dados do usuário.
- **Permissões e Segurança:** Controle de acesso por número de telefone e permissões de owner.
- **Respostas Personalizadas:** Mensagens customizadas para cada etapa e situação.
- **Deploy em Kubernetes:** Pronto para ser executado em ambientes de produção com arquivos de deployment e configuração.

---

## 🛠️ Tecnologias Utilizadas

- **Go (Golang):** Linguagem principal do projeto.
- **WhatsApp API:** Integração para envio e recebimento de mensagens.
- **SQLite:** Persistência de dados dos usuários e estágios.
- **Kubernetes:** Orquestração e deploy.
- **GitHub Actions:** CI/CD, testes, lint e segurança.

---


