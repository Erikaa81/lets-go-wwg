using System;

namespace DIO.Series
{
    public class Serie : EntidadeBase
    {
        public int Id { get; private set; }

        // Atributos
        private Genero Genero { get; set; }
		private string Titulo { get; set; }
		private string Descricao { get; set; }
		private int Ano { get; set; }
		private string Classificacao { get; set; }
		private int Avaliacao {get; set;}
        private bool Excluido {get; set;}

        // Métodos
		public Serie(int id, Genero genero, string titulo, string descricao, int ano, string classificacao, int avaliacao)
		{
			this.Id = id;
			this.Genero = genero;
			this.Titulo = titulo;
			this.Descricao = descricao;
			this.Ano = ano;
			this.Classificacao = classificacao;
			this.Avaliacao = avaliacao;
            this.Excluido = false;
		}

        public override string ToString()
		{
			// Environment.NewLine https://docs.microsoft.com/en-us/dotnet/api/system.environment.newline?view=netcore-3.1
            string retorno = "";
            retorno += "Gênero: " + this.Genero + Environment.NewLine;
            retorno += "Titulo: " + this.Titulo + Environment.NewLine;
            retorno += "Descrição: " + this.Descricao + Environment.NewLine;
            retorno += "Ano de Início: " + this.Ano + Environment.NewLine;
			retorno += "Classificação Indicativa: " + this.Classificacao + Environment.NewLine;
			retorno += "Avaliação: " + this.Avaliacao + Environment.NewLine;
            retorno += "Excluido: " + this.Excluido;
			return retorno;
		}

        public string retornaTitulo()
		{
			return this.Titulo;
		}

		public int retornaId()
		{
			return this.Id;
		}
        public bool retornaExcluido()
		{
			return this.Excluido;
		}
        public void Excluir() {
            this.Excluido = true;
        }
    }

    public class EntidadeBase
    {
    }
}